import { Breadcrumb, BreadcrumbItem, BreadcrumbLink, BreadcrumbList, BreadcrumbPage, BreadcrumbSeparator } from "@/components/ui/breadcrumb";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Tabs, TabsList, TabsContent, TabsTrigger } from "@/components/ui/tabs";
import { Link } from "react-router-dom";
import { PATH } from "@/routes/routes";

export default function ReleasesPage() {
	return (
		<>
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
							<Link to={PATH.USER.PROJECT}>KiCad: Universal Interface</Link>
					</BreadcrumbItem>
					<BreadcrumbSeparator />
					<BreadcrumbItem>
						<BreadcrumbPage>
							<Link to="#">Project Releases</Link>
						</BreadcrumbPage>
					</BreadcrumbItem>
				</BreadcrumbList>
			</Breadcrumb>
			<div className="h-[80vh]">
				{/* <script type="module" src="@/assets/kicanvas.js"></script> */}
				<Tabs defaultValue="ovw">
					<div className="flex">
						<Select defaultValue="main" >
							<SelectTrigger className="w-[150px]">
								<div className="text-muted-foreground">Release: </div>
								<SelectValue />
							</SelectTrigger>
							<SelectContent>
								<SelectItem value="main">v1.0</SelectItem>
								<SelectItem value="dev">v1.1</SelectItem>
								{/* TODO: Make responsive to selected project */}
							</SelectContent>
						</Select>
						<div className="px-2"></div>
						<TabsList>
							<TabsTrigger value="ovw">Overview</TabsTrigger>
							<TabsTrigger value="sch">Schematic</TabsTrigger>
							<TabsTrigger value="pcb">Board</TabsTrigger>
							<TabsTrigger value="3dv">3D</TabsTrigger> {/* NOTE: Stands for 3D View */}
							<TabsTrigger value="drt">Drafting</TabsTrigger>
							<TabsTrigger value="grb">Gerber</TabsTrigger>
							<TabsTrigger value="bom">BOM</TabsTrigger>
							<TabsTrigger value="amb">Assembly</TabsTrigger>
						</TabsList>
					</div>
					<TabsContent forceMount value="ovw" className='data-[state=inactive]:hidden'>
						<div className="h-[75vh]">
							<div className="flex flex-col min-h-full justify-center px-20">
								<div>
									Project releases page coming soon ⚗️📖😎
								</div>
								<div className="py-2"></div>
								<ol>
									<li>This view will have the following features:</li>
									<li>Schematic view</li>
									<li>Board view</li>
									<li>3D view</li>
									<li>Drawing view</li>
									<li>BOM view</li>
									<li>Assembly view</li>
								</ol>
								<div className="grow-2 p-28"></div>
							</div>
						</div>
					</TabsContent>
				</Tabs>
			</div>
		</>
	)
}