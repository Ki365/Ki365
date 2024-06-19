import { Breadcrumb, BreadcrumbItem, BreadcrumbLink, BreadcrumbList, BreadcrumbPage, BreadcrumbSeparator } from "@/components/ui/breadcrumb";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Tabs, TabsList, TabsContent, TabsTrigger } from "@/components/ui/tabs";
import DashboardTemplate from "@/templates/dashboard-template";
import { useEffect } from "react";
import { Link } from "react-router-dom";

// import script from '@/assets/kicanvas.js'

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

export default function ProjectPage() {
    useScript("/kicanvas.js")

    return (
        <>
            {/* <div className="flex items-center">
                <h1 className="text-lg font-semibold md:text-2x1">Olimex: A64-OLinuXino</h1>
            </div> */}
            <Breadcrumb className="">
                <BreadcrumbList>
                    <BreadcrumbItem>
                        <BreadcrumbLink>
                            <Link to="/dashboard/">Dashboard</Link>
                        </BreadcrumbLink>
                    </BreadcrumbItem>
                    <BreadcrumbSeparator />
                    <BreadcrumbItem>
                        <BreadcrumbLink>
                            <Link to="/dashboard/projects">Projects</Link>
                        </BreadcrumbLink>
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
                {/* TODO: Get theme selector to work by default */}
                <Tabs defaultValue="sch">
                    <div className="flex">
                        <Select>
                            <SelectTrigger className="w-[150px]">
                                <div className="text-muted-foreground">Branch: </div>
                                <SelectValue defaultValue="main" />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectItem value="main">main</SelectItem>
                                <SelectItem value="dev">dev</SelectItem>
                                {/* Make responsive to selected project */}
                            </SelectContent>
                        </Select>
                        <div className="px-2"></div>
                        <TabsList>
                            <TabsTrigger value="sch">Schematic</TabsTrigger>
                            <TabsTrigger value="pcb">Board</TabsTrigger>
                            <TabsTrigger value="3dv">3D</TabsTrigger> {/* NOTE: Stands for 3D View */}
                            <TabsTrigger value="bom">BOM</TabsTrigger>
                        </TabsList>
                    </div>
                    <TabsContent value="sch">
                        <div className="h-[75vh]">
                            <kicanvas-embed src="/interf_u/interf_u.kicad_sch" theme="kicad" controls="full" />
                        </div>
                    </TabsContent>
                    <TabsContent value="pcb">
                        <div className="h-[75vh]">
                            <kicanvas-embed src="/interf_u/interf_u.kicad_pcb" theme="kicad" controls="full" />
                        </div>
                    </TabsContent>
                </Tabs>
            </div>
        </>
    )
}