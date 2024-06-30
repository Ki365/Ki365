import { ImportIcon } from "lucide-react";
import { Button } from "../ui/button";
import { toast } from "sonner";
import { Pagination, PaginationContent, PaginationEllipsis, PaginationItem, PaginationLink, PaginationNext, PaginationPrevious } from "../ui/pagination";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "../ui/card";
import Project from "../dashboard/project";

import RingLoader from "react-spinners/RingLoader";
import { useEffect, useState } from "react";

interface Project {
	ID: number
	Image: String
	ProjectName: String
	Description: String
}

function RenderProjects(props : any) {
	if (props.projects) {
		return (
			<>
				{props.projects.map((project : Project, index: number) => (
					<Project
						key={index}
						_id={project.ID}
						image={project.Image}
						projectName={project.ProjectName}
						description={project.Description}
					/>
				))}
			</>
		)
	} else {
		// TODO: remove error in favor of better error handling
		console.error("Error mapping project prop as project card")
	}
}

function ContentEmpty() {
    return (
        <div className="flex flex-1 items-center justify-center rounded-lg border border-dashed shadow-sm" x-chunk="dashboard-02-chunk-1">
            <div className="flex flex-col items-center gap-1 text-center">
                <ImportIcon size={100} strokeWidth={1} className="text-muted-foreground m-4" />
                <h3 className="text-2x1 font-bold tracking-tight">
                    You have no projects yet
                </h3>
                <p className="text-sm text-muted-foreground">
                    You can start collaborating as soon as you add a project.
                </p>
                <div className="p-3"></div>
                <Button className="mt-4" onClick={() => {
                    const promise = () => new Promise((resolve) => setTimeout(() => resolve({ name: 'Example' }), 3000));

                    toast.promise(promise, {
                        loading: 'Adding example project, please wait...',
                        success: (data: any) => {
                            return `${data.name} project has been added!`;
                        },
                        error: 'Error',
                        duration: 5000,
                        action: {
                            label: 'Cancel',
                            onClick: () => console.log('Action!'),
                        },
                    })
                }}>
                    Add an example project</Button>
                <Button className="mt-4">Add a project</Button>
            </div>
        </div>
    )
}

function ContentFull(props : any) {
    return (
        <div className="flex flex-1 items-center justify-center rounded-lg border shadow-sm" x-chunk="dashboard-02-chunk-1">
            <div className="container my-6s">
                <div className="grid md:grid-cols-2 sm:grid-cols-1 lg:grid-cols-4 gap-4">
						<RenderProjects projects={props.projects}/>
                    <div className="flex">
                        <Card x-chunk="dashboard-02-chunk-0" className="p-4 flex flex-col justify-between bg-[linear-gradient(to_bottom_right,var(--cyan-500),var(--emerald-500))]"> {/* bg-[linear-gradient(to_bottom_right,var(--cyan-500),var(--emerald-500))] */}
                            <CardHeader className="p-2 pt-0 md:p-4 text-white">
                                <CardTitle>Visit the docs!</CardTitle>
                                <CardDescription className="text-white py-2">
                                    Learn about all the features Ki365 has to offer. From project collaboration to discovery, Ki365 is your go-to place for engineering content!
                                </CardDescription>
                            </CardHeader>
                            <CardContent className="p-2 pt-0 md:p-4 md:pt-0 ">
                                <Button size="sm" className="w-full">
                                    <a href="https://ki365.github.io/docs/">
                                        Visit the docs
                                    </a>
                                </Button>
                            </CardContent>
                        </Card>
                    </div>
                </div>
                <div className="p-2"></div>
                {/* <Pagination disabled="true"> */}
                <Pagination>
                    <PaginationContent>
                        <PaginationItem>
                            <PaginationPrevious href="#" />
                        </PaginationItem>
                        <PaginationItem>
                            <PaginationLink href="#">1</PaginationLink>
                        </PaginationItem>
                        <PaginationItem>
                            <PaginationLink href="#" isActive>
                                2
                            </PaginationLink>
                        </PaginationItem>
                        <PaginationItem>
                            <PaginationLink href="#">3</PaginationLink>
                        </PaginationItem>
                        <PaginationItem>
                            <PaginationEllipsis />
                        </PaginationItem>
                        <PaginationItem>
                            <PaginationNext href="#" />
                        </PaginationItem>
                    </PaginationContent>
                </Pagination>
            </div>
        </div>
    )
}


export default function DashboardContnet() {
    const [hasResolved, setHasResolved] = useState(false)
    const [hasErrors, setHasErrors] = useState(false)
    const [hasProjects, setHasProjects] = useState(false)
    const [projects, setProjects] = useState([])
	
	// TODO: Add better user visibility of errors
	// TODO: Add timeout to promise
	useEffect(() => {
		new Promise<any>(() => 
			fetch("/api/projects", {
				method: "GET",
			}).then((data) => {
				console.log(data)
				if (data?.ok) {
					if(data?.status != 204) {
						data.json().then((data) => {
							console.log(data)
							console.log("success")
							setProjects(data.projects)
							setHasResolved(true)
							setHasErrors(false)
							setHasProjects(true)
						}).catch(error => {
							console.log("Error parsing project JSON data:")
							console.log(error)
							setHasResolved(true)
							setHasErrors(true)
						})
					} else {
						setHasResolved(true)
						setHasErrors(false)
						setHasProjects(false)
						console.log("no projects")
					} 
				} else {
					setHasResolved(true)
					setHasErrors(true)
				}
			}
		).catch(error => {
				console.log(error)
				setHasErrors(true)
			})
		)
	}, [])
		

	if(hasResolved) {
		if(!hasErrors) {
			if(hasProjects) {
				return <ContentFull projects={projects}/>
			} else {
				return <ContentEmpty />
			}
		} else {
			return(
			<>
				<div className="flex flex-col justify-center items-center gap-4">
					<p>Errors!!!</p>
				</div>
			</>
			)
		}
	} else {
		return (
			<>
				<div className="flex flex-col justify-center items-center gap-4">
					<RingLoader />
					<p>Loading content...</p>
				</div>
			</>
		)
	}
}