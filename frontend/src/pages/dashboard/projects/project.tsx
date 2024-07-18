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
import { Table, TableBody, TableCell, TableFooter, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import RingLoader from "react-spinners/RingLoader";
import { ScrollArea } from "@/components/ui/scroll-area";

// TODO: make responsive
const bom = [
    {
        Reference: "C1",
        Value: "10n",
        Datasheet: "",
        Footprint: "SMD-CMN:SMD-0603C",
        Qty: "1",
        DNP: "",
        Num: "1",
        Description: ""
    },
    {
      Reference: "C1",
      Value: "10n",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "1",
      Description: ""
    },
    {
      Reference: "C2",
      Value: "CAPACITOR",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "2",
      Description: ""
    },
    {
      Reference: "C3",
      Value: "CAPACITOR",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "3",
      Description: ""
    },
    {
      Reference: "C4",
      Value: "20pF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "4",
      Description: ""
    },
    {
      Reference: "C5",
      Value: "20pF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "5",
      Description: ""
    },
    {
      Reference: "C6",
      Value: "0.1uF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "6",
      Description: ""
    },
    {
      Reference: "C7",
      Value: "0.1uF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "7",
      Description: ""
    },
    {
      Reference: "C8",
      Value: "0.1uF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "8",
      Description: ""
    },
    {
      Reference: "C9",
      Value: "0.1uF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "9",
      Description: ""
    },
    {
      Reference: "C10",
      Value: "0.1uF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "10",
      Description: ""
    },
    {
      Reference: "C11",
      Value: "10nF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "11",
      Description: ""
    },
    {
      Reference: "C12",
      Value: "10nF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "12",
      Description: ""
    },
    {
      Reference: "C13",
      Value: "2.2uF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "13",
      Description: ""
    },
    {
      Reference: "C14",
      Value: "2.2uF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "14",
      Description: ""
    },
    {
      Reference: "C15",
      Value: "10uF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-1206C",
      Qty: "1",
      DNP: "",
      Num: "15",
      Description: ""
    },
    {
      Reference: "C16",
      Value: "4.7uF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0805C",
      Qty: "1",
      DNP: "",
      Num: "16",
      Description: ""
    },
    {
      Reference: "C17",
      Value: "10uF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-1206C",
      Qty: "1",
      DNP: "",
      Num: "17",
      Description: ""
    },
    {
      Reference: "C18",
      Value: "1uF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "18",
      Description: ""
    },
    {
      Reference: "C19",
      Value: "0.1uF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "19",
      Description: ""
    },
    {
      Reference: "C20",
      Value: "0.1uF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "20",
      Description: ""
    },
    {
      Reference: "C21",
      Value: "1uF",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "21",
      Description: ""
    },
    {
      Reference: "C22",
      Value: "100n",
      Datasheet: "",
      Footprint: "SMD-CMN:SMD-0603C",
      Qty: "1",
      DNP: "",
      Num: "22",
      Description: ""
    },
  ]

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
                        <div className="flex flex-col justify-between h-[75vh]">
                            <ScrollArea className="rounded-md border h-max-[70vh]" >
                                <Table>
                                    <TableHeader>
                                        <TableRow className="divide-x">
                                            <TableHead>Designator</TableHead>
                                            <TableHead>Value</TableHead>
                                            <TableHead>Description</TableHead>
                                            <TableHead>Footprint</TableHead>
                                            <TableHead>DNP</TableHead>
                                            <TableHead className=" text-right">Quantity</TableHead>
                                        </TableRow>
                                    </TableHeader>
                                    {/* <div className="static"> */}
                                    <TableBody className="">
                                        {bom.map(b => (
                                            <TableRow key={b.Reference} className="divide-x">
                                                <TableCell className="font-medium">{b.Reference}</TableCell>
                                                <TableCell className="">{b.Value}</TableCell>
                                                <TableCell className="">{b.Description}</TableCell>
                                                <TableCell className="">{b.Footprint}</TableCell>
                                                <TableCell className="">{b.DNP}</TableCell>
                                                <TableCell className="text-right">{b.Qty}</TableCell>
                                            </TableRow>
                                        ))}
                                    </TableBody>
                                    <TableFooter>
                                        <TableRow>
                                            <TableCell colSpan={5}>Total</TableCell>
                                            <TableCell className="text-right">{bom.reduce((acc, curr) => acc + ~~Number(curr.Qty), 0)}:{bom.length}</TableCell>
                                        </TableRow>
                                    </TableFooter>
                                </Table>
                            </ScrollArea>
                            <div className="flex flex-col justify-center items-center m-3">BOM data for building: 🤯</div>
                        </div>
                    </TabsContent>
                </Tabs>
            </div>
        </>
    )
}