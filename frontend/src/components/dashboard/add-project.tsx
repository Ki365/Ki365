import { AlertDialog, AlertDialogAction, AlertDialogCancel, AlertDialogContent, AlertDialogDescription, AlertDialogFooter, AlertDialogHeader, AlertDialogTitle, AlertDialogTrigger } from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";
import { PlusIcon } from "lucide-react";
import { useRef, useState } from "react";


export function AddProject() {

	const file = useRef<HTMLInputElement>(document.createElement("input"))
	const ref = useRef<HTMLFormElement>(document.createElement("form"))

	const [isDialogOpen, setIsDialogOpen] = useState(false);
	const [imageURL, setImageURL] = useState("");

	const handleClick = () => {
		file.current?.click()
	}
	const handleFileAdd = (event: any) => {
		let value = event.target.files[0].name
		setIsDialogOpen(true)
		setImageURL(value)
	}
	const handleConfirm = () => {
		ref.current?.requestSubmit()
		setIsDialogOpen(false)
	}

	const handleSubmit = (event: any) => {
		event.preventDefault()
		let form = event.target
		let errors = false;
		// TODO: local error checking here
		if (!errors) {
			sendData(form)
		}
	}


	async function sendData(form: any) {
		const formData = new FormData(form)

		try {
			const response = await fetch('/api/upload/project', {
				method: "POST",
				body: formData
			})
			console.log(await response)
		} catch (e) {
			console.log(e)
		}
	}

	return (
		<>
			<form
				encType="multipart/form-data"
				action="/api/upload/project"
				method="post"
				ref={ref}
				className="hidden"
				onChange={handleFileAdd}
				onSubmit={handleSubmit}
			>
				<input type="file" ref={file} name="project-archive" />
			</form>
			<div className="px-2">
				<Button onClick={handleClick}>
					<PlusIcon size={20} />
				</Button>
			</div>

			<AlertDialog open={isDialogOpen}>
				<AlertDialogContent>
					<AlertDialogHeader>
						<AlertDialogTitle>Are you sure you want to upload this file:</AlertDialogTitle>
						<AlertDialogDescription>
							{imageURL}
						</AlertDialogDescription>
					</AlertDialogHeader>
					<AlertDialogFooter>
						<AlertDialogCancel onClick={() => setIsDialogOpen(false)}>Cancel</AlertDialogCancel>
						<AlertDialogAction onClick={handleConfirm}>Continue</AlertDialogAction>
					</AlertDialogFooter>
				</AlertDialogContent>
			</AlertDialog>
		</>
	)
}
