import { AlertDialog, AlertDialogAction, AlertDialogCancel, AlertDialogContent, AlertDialogDescription, AlertDialogFooter, AlertDialogHeader, AlertDialogTitle } from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";
import { PlusIcon } from "lucide-react";
import { FormEvent, useRef, useState } from "react";
import { toast } from "sonner";
import { Checkbox } from "@/components/ui/checkbox";
import { Dialog, DialogContent, DialogTitle, DialogDescription } from "@/components/ui/dialog";
import { DialogHeader, DialogFooter } from "@/components/ui/dialog";

export function AddProject() {

	const file = useRef<HTMLInputElement>(document.createElement("input"))
	const ref = useRef<HTMLFormElement>(document.createElement("form"))

	const [isConfirmationDialogOpen, setIsConfirmationDialogOpen] = useState(false);
	const [isWarningDialogOpen, setIsWarningDialogOpen] = useState(false);
	const [imageURL, setImageURL] = useState("");

	const handleAddProjectButton = () => {
		setIsWarningDialogOpen(true)
	}
	const handleClick = () => {
		file.current?.click()
	}
	const handleFileAdd = (event: any) => {
		let value = event.target.files[0].name
		setIsConfirmationDialogOpen(true)
		setImageURL(value)
	}
	const handleConfirm = () => {
		ref.current?.requestSubmit()
		setIsConfirmationDialogOpen(false)
	}

	const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
		event.preventDefault()
		let form = event.currentTarget
		let errors = false;
		// TODO: local error checking here
		if (!errors) {
			sendData(form)
		}
	}

	const cancelRequest = new AbortController()

	async function sendData(form: HTMLFormElement) {
		const formData = new FormData(form)

		try {
			const response = new Promise<any>((resolve, reject) =>
				fetch('/api/upload/project', {
					method: "POST",
					body: formData,
					signal: cancelRequest.signal
				})
					.then((data) => {
						if (data?.ok) {
							resolve(1)
						} else {
							reject()
						}
					})
					.catch(error => {
						console.error(error)
						reject()
					}).finally(() => {
						form.reset()
					}))

			toaster(response)

		} catch (e) {
			console.log(e)
		}
	}

	const toaster = (promise: Promise<any>) => {
		toast.promise(promise, {
			loading: 'Adding project named: , please wait...',
			success: (data: any) => {
				return `${data.name} project has been added!`;
			},
			// TODO: Dynamically change error msg (PromiseTResult) based on error code
			error: "Error",
			duration: 10000,
			cancel: {
				// TODO: Dynamically change label text for "cancel" or "dimiss" operation and change onClick function calls
				label: 'Cancel',
				onClick: (event) => {
					event.preventDefault()
					cancelRequest.abort()
					console.log("Cancelled file upload!")

				},
			},
		})
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
				<input type="file" ref={file} name="project-archive" accept=".zip,.tar" />
			</form>
			<div className="px-2">
				<Button onClick={handleAddProjectButton}>
					<PlusIcon size={20} />
				</Button>
			</div>

			<AlertDialog open={isConfirmationDialogOpen}>
				<AlertDialogContent>
					<AlertDialogHeader>
						<AlertDialogTitle>Are you sure you want to upload this file:</AlertDialogTitle>
						<AlertDialogDescription>
							{imageURL}
						</AlertDialogDescription>
					</AlertDialogHeader>
					<AlertDialogFooter>
						<AlertDialogCancel onClick={() => setIsConfirmationDialogOpen(false)}>Cancel</AlertDialogCancel>
						<AlertDialogAction onClick={handleConfirm}>Continue</AlertDialogAction>
					</AlertDialogFooter>
				</AlertDialogContent>
			</AlertDialog>

			<Dialog open={isWarningDialogOpen} onOpenChange={() => (setIsWarningDialogOpen(false))}>
				<DialogContent className="sm:max-w-[650px]">
					<DialogHeader>
						<DialogTitle>Please read the following information regarding importing projects:</DialogTitle>
						<DialogDescription>
							When you are finished reading the following notes, click confirm.
						</DialogDescription>
					</DialogHeader>
					<div className="flex flex-col space-y-10">
						<div >
							<ul className="space-y-2">
								<li>Ki365 is currently in a pre-production <b><i>pre-alpha</i></b> state. This means any imported project data may be corrupted at any time and/or for any reason.</li>
								<li>This Ki365 instance has been configured with internal storage. Project features (including Export) <b><i>may not work correctly</i></b> or at all.</li>
								<li>Project Import currently only supports zipped files of KiCad projects (including git repository support). Increased functionality will be added at a later date.</li>
								<li>If a zipped git repository is used, please ensure all files are committed to a branch before uploading. Ki365 does not currently support native writing to repositories.</li>
								<li>Finally, the use of this program is intended only for <b><i>evaluation purposes only.</i></b></li>

							</ul>
						</div>
					<div className="items-top flex space-x-2">
						<Checkbox id="terms1" />
						<div className="grid gap-1.5 leading-none">
							<label
								htmlFor="terms1"
								className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
							>
								I confirm I have read and acknoledged the warnings listed above.
								I understand any imported project data may be lost partially or in its entirety.
								I assume complete liability for any loss of data or personal/corporate damages caused by this program during its pre-production status.
								I acknoledge no support is guarenteed, except bona fide, for this program or by its representatives at this time. 
							</label>
							<p className="text-sm text-muted-foreground">
								To disable this dialog, go to Preferences {'>'} Warnings {'>'} Project Import Warning
							</p>
						</div>
					</div>
					</div>
					<DialogFooter>
						<Button type="submit" onClick={() => {handleClick(); setIsWarningDialogOpen(false)}}>Confirm</Button>
					</DialogFooter>
				</DialogContent>
			</Dialog>
		</>
	)
}
