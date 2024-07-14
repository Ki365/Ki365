import { Breadcrumb, BreadcrumbItem, BreadcrumbList, BreadcrumbPage, BreadcrumbSeparator } from "@/components/ui/breadcrumb";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Tabs, TabsList, TabsContent, TabsTrigger } from "@/components/ui/tabs";
import React, { Suspense, useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { Canvas } from "@react-three/fiber"
import { PerspectiveCamera, Stage, useGLTF } from '@react-three/drei'
import { OrbitControls } from '@react-three/drei'
import { PATH } from "@/routes/routes";
import { Button } from "@/components/ui/button";
import RingLoader from "react-spinners/RingLoader";

declare global {
    namespace JSX {
      interface IntrinsicElements {
        ['kicanvas-embed']: React.DetailedHTMLProps<React.HTMLAttributes<HTMLElement>, HTMLElement> & {
			fileType: React.ReactNode;
            src: React.ReactNode;
            theme: React.ReactNode;
            controls: React.ReactNode;
			controlslist: React.ReactNode;
        };
      }
    }
  }

const useScript: any = (url: any) => {
    useEffect(() => {
        const script = document.createElement('script');

        script.src = url;
        script.async = true;

        document.body.appendChild(script);

        return () => {
            document.body.removeChild(script);
        }
    }, [url]);
};

const RenderCanvas = (props: {type : string, id : any}) => {
	// const [resources, setResources] = useState<File[]>([])
	// const [resourcesPath, setResourcesPath] = useState<String[]>([])
	const [resources, setResources] = useState<File>()
	const [resourcesPath, setResourcesPath] = useState<String>()
	const [startCanvas, setStartCanvas] = useState(false)

	useEffect(() => {
		var apiType : string
		if (props.type == "sch") {
			apiType = "schematics"
		} else if (props.type == "pcb") {
			apiType = "layouts"
		} else if (props.type == "pro") {
			apiType = "project"
		} else if (props.type == "mod") {
			apiType = "models"
		} else {
            return
        }
		
		fetch(`/api/projects/${props.id}/${apiType}`)
		.then((response) => {
			return response.blob()
		}).then((blob) => {
			// setResource(existing => [...existing, (new File([blob], "ii.kicad_sch"))])
			setResources(new File([blob], `${props.id}-${apiType}-temp`, {
				type: blob.type,
			}))
		})
	}, [])

	useEffect(() => {
		if(resources) {
			setResourcesPath(URL.createObjectURL(resources))
			setStartCanvas(true)
		}
	}, [resources])

	if (startCanvas) {
        if (props.type == "mod") {
            if (resourcesPath != null) {
                const gltf = useGLTF(`${resourcesPath}`)

                return (
                    <>
                        <Suspense>
                            <Canvas className="max-h-fit  sm:aspect-[2/2] md:aspect-[2/1.85] max-md:[2/1.75] lg:aspect-[2/1.5] max-lg:aspect-auto aspect-[3/2]" style={{ background: '#dfdfdf' }}>
                                <PerspectiveCamera makeDefault position={[0, 1, 5]} />
                                <ambientLight />
                                <Stage preset="rembrandt" intensity={0.01} environment="warehouse">
                                    <primitive
                                        object={gltf.scene}
                                        />
                                </Stage>
                                <OrbitControls />
                            </Canvas>
                        </Suspense>
                    </>
                )
            }
        }
		return (
			<>
				<kicanvas-embed fileType={props.type} src={resourcesPath} theme="kicad" controls="full" controlslist="nooverlay"/>
			</>
		)
	} else {
        if (props.type == "mod") {
            return (
                <>
                    <div className="flex flex-col items-center justify-center content-center h-full gap-8">
                        <RingLoader />
                        <div>Loading Model viewer...</div>
                    </div>
                </>
            )
        }
		return (
			<>
				<div className="flex flex-col items-center justify-center content-center h-full gap-8">
					<RingLoader />
					<div>Loading schematic...</div>
				</div>
			</>
		)
	}
}

export default function ProjectPage() {
	const { id } = useParams()

    useScript("/kicanvas.js")

    return (
        <>
            {/* <div className="flex items-center">
                <h1 className="text-lg font-semibold md:text-2x1">Olimex: A64-OLinuXino</h1>
            </div> */}
            <Breadcrumb className="">
                <BreadcrumbList>
                    <BreadcrumbItem>
                            <Link to={PATH.USER.HOME}>Dashboard</Link>
                    </BreadcrumbItem>
                    <BreadcrumbSeparator />
                    <BreadcrumbItem>
                            <Link to={PATH.USER.ALL_PROJECTS}>Projects</Link>
                    </BreadcrumbItem>
                    <BreadcrumbSeparator />
                    <BreadcrumbItem>
                        <BreadcrumbPage>
                            <Link to="#">KiCad: Universal Interface</Link>
                        </BreadcrumbPage>
                    </BreadcrumbItem>
                </BreadcrumbList>
            </Breadcrumb>
            <div className="h-[80vh]">
                {/* <script type="module" src="@/assets/kicanvas.js"></script> */}
                <Tabs defaultValue="sch">
                    <div className="flex flex-wrap gap-x-4 gap-y-2">
                        <Select defaultValue="main" >
                            <SelectTrigger className="w-[150px]">
                                <div className="text-muted-foreground">Branch: </div>
                                <SelectValue />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectItem value="main">main</SelectItem>
                                <SelectItem value="dev">dev</SelectItem>
                                {/* TODO: Make responsive to selected project */}
                                {/* TODO: Limit KiCanvas to just this page*/}
                            </SelectContent>
                        </Select>
                        <TabsList>
                            <TabsTrigger value="sch">Schematic</TabsTrigger>
                            <TabsTrigger value="pcb">Board</TabsTrigger>
                            <TabsTrigger value="3dv">3D</TabsTrigger> {/* NOTE: Stands for 3D View */}
                            <TabsTrigger value="drt">Drafting</TabsTrigger> {/* NOTE: Stands for 3D View */}
                            <TabsTrigger value="bom">BOM</TabsTrigger>
                        </TabsList>
						<Link to={PATH.USER.PROJECT_RELEASES}>
							<Button>View Releases</Button>
						</Link>
                    </div>
                    <TabsContent forceMount value="sch" className='data-[state=inactive]:hidden'>
                        <div className="h-[75vh]">
                            <RenderCanvas id={id} type="sch"/>
                        </div>
                    </TabsContent>
                    <TabsContent forceMount value="pcb" className='data-[state=inactive]:hidden'>
                        <div className="h-[75vh]">
							<RenderCanvas id={id} type="pcb"/>
                        </div>
                    </TabsContent>
                    <TabsContent forceMount value="3dv" className='data-[state=inactive]:hidden'>
                        <div className="h-[75vh]">
                            <RenderCanvas id={id} type="mod"/>
                        </div>
                    </TabsContent>
					<TabsContent forceMount value="drt" className='data-[state=inactive]:hidden'>
                        <div className="h-[75vh]">
                            <div className="flex flex-col min-h-full justify-center items-center">
                                <div>
                                    Drafting View requires additional changes to the KiCad source code.
                                </div>
								<div className="py-1"></div>
								<div>
									To learn more visit this&nbsp;
									<a href="https://gitlab.com/kicad/code/kicad/-/issues/2276"
									target="_blank"
									rel="noreferrer noopener">
										<b>Link</b>
									</a>
									.
								</div>
                                <div className="grow-2 p-28"></div>
                            </div>
                        </div>
                    </TabsContent>
                    <TabsContent forceMount value="bom" className='data-[state=inactive]:hidden'>
                        <div className="h-[75vh]">
                            <div className="flex flex-col min-h-full justify-center items-center">
                                <div>
                                    BOM data coming soon 🧨💥🤯
                                </div>
                                <div className="grow-2 p-28"></div>
                            </div>
                        </div>
                    </TabsContent>
                </Tabs>
            </div>
        </>
    )
}