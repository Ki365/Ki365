import { AlertDialog, AlertDialogAction, AlertDialogCancel, AlertDialogContent, AlertDialogDescription, AlertDialogFooter, AlertDialogHeader, AlertDialogTitle } from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";
import { PlusIcon } from "lucide-react";
import { FormEvent, useContext, useEffect, useRef, useState } from "react";
import { toast } from "sonner";
import { Checkbox } from "@/components/ui/checkbox";
import { Dialog, DialogContent, DialogTitle, DialogDescription, DialogClose } from "@/components/ui/dialog";
import { DialogHeader, DialogFooter } from "@/components/ui/dialog";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuSeparator, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { UserContext } from '@/pages/dashboard/projects/all-projects'

export function AddProject() {

	const file = useRef<HTMLInputElement>(document.createElement("input"))
	const ref = useRef<HTMLFormElement>(document.createElement("form"))

	const [isConfirmationDialogOpen, setIsConfirmationDialogOpen] = useState(false);
	const [isWarningDialogOpen, setIsWarningDialogOpen] = useState(false);
	const [imageURL, setImageURL] = useState("");

	const [isLinkOpen, setIsLinkOpen] = useState(false);

	const [isEnableExampleProjects, setEnableExampleProjects] = useState(false);

	const [hasResolvedExampleProjectsResponse, setHasResolvedExampleProjectsResponse] = useState(false);

	const  context = useContext(UserContext)
	
	useEffect(() => {
		new Promise<any>(() =>
			fetch('/api/toggle-example-projects', {
				method: "GET",
				headers: {
					Accept: 'application/json',
				},
				// signal: cancelRequest.signal
			}).then((data) => {
				// console.log(data)
				return data.json()
			}).then((data) => {
				console.log(data)
				if (data == true) {
					setEnableExampleProjects(true)
					setHasResolvedExampleProjectsResponse(true)
				} else if (data == false) {
					setEnableExampleProjects(false)
					setHasResolvedExampleProjectsResponse(true)
				} else {
					setEnableExampleProjects(false)
					setHasResolvedExampleProjectsResponse(false)
				}
			}).catch(error => {
				console.error(error)
				// reject()
			}))
	}, [])

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

	const handleToggleExampleProjects = () => {
		setEnableExampleProjects(!isEnableExampleProjects)
		var method : string
		if (isEnableExampleProjects) {
			method = "DELETE"
		} else {
			method = "POST"
		}
		try {
			const response = new Promise<any>((resolve, reject) =>
				fetch('/api/toggle-example-projects', {
					method: method,
					body: "",
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
						// props.setShouldResolve(true)
						context?.setUser(true)
					}))

			// TODO: make toaster component unique to this toggle
			toaster(response)

		} catch (e) {
			console.log(e)
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
				<DropdownMenu modal={false}>
					<DropdownMenuTrigger asChild>
						<Button variant="secondary" className="rounded" title='Add project'>
							<PlusIcon size={20} className="h-6 w-6" />
							<span className="sr-only">Toggle user menu</span>
						</Button>
					</DropdownMenuTrigger>
					<DropdownMenuContent align="end">
						<DropdownMenuItem onClick={() => setIsLinkOpen(true)}>Add Link</DropdownMenuItem>
						<DropdownMenuItem>Add Mirror</DropdownMenuItem>
						<DropdownMenuItem>Add Local</DropdownMenuItem>
						<DropdownMenuSeparator />
						<DropdownMenuItem onClick={handleAddProjectButton}>Import Archive</DropdownMenuItem>
						<DropdownMenuSeparator />
						<DropdownMenuItem 
							onClick={handleToggleExampleProjects}
							disabled={!hasResolvedExampleProjectsResponse}
							>
								Toggle example projects
						</DropdownMenuItem>
					</DropdownMenuContent>
				</DropdownMenu>
			</div>

			<Dialog open={isLinkOpen} onOpenChange={() => (setIsLinkOpen(false))}>
				<DialogContent>
					<DialogHeader>
						<DialogTitle>Add a project as a Linked Connection</DialogTitle>
						<DialogDescription>
						Linked repositories allow data owners to submit changes to Ki365 and have them forwarded to the linked data store as a form of offsite backup.
						</DialogDescription>
					</DialogHeader>
					<div className="grid flex-1 gap-2">
						<Label htmlFor="link" className="sr-only">Link</Label>
						<Input id="link" placeholder="https://github.io/Ki365/example-project" />
					</div>
					<DialogFooter className="sm:justify-start">
						<div className="grid flex-1 gap-2">
							<Label className="p-2">Coming soon: Authentication for private repos!</Label>
							<div className="flex flex-row justify-between w-full p-">
								<DialogClose asChild>
									<Button type="button" variant="secondary">Close</Button>
								</DialogClose>
								{/* TODO: Add functionality to this button */}
								<Button type="button" variant="default">Submit</Button>
							</div>
						</div>
					</DialogFooter>
				</DialogContent>
			</Dialog>

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
