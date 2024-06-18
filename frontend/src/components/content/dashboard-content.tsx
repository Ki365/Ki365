import { ImportIcon } from "lucide-react";
import { Button } from "../ui/button";
import { toast } from "sonner";
import { Pagination, PaginationContent, PaginationEllipsis, PaginationItem, PaginationLink, PaginationNext, PaginationPrevious } from "../ui/pagination";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "../ui/card";
import Project from "../dashboard/project";

const projects = [
    {
        _id: "100001",
        image: "https://www.kicad.org/made-with-kicad/olinuxino/olinuxino.jpg",
        projectName: "Olimex: A64-OLinuXino",
        // description: "single-board computer running Linux & Android. The design based on a 64-bit ARM CPU and includes 1 or 2 GB DDR3 RAM, 4 GB flash memory, microSD card socket, WiFi & BLE4.0, Ethernet, HDMI & audio output.",
        description: "single-board computer running Linux & ...",
    },
    {
        _id: "100001",
        image: "https://www.kicad.org/made-with-kicad/ciaa-project-acc/CIAA-ACC.jpg",
        projectName: "CIAA: CIAA-ACC (HPC)",
        // description: "12 layer PCB design for High Performance Computing. It has a Xilinx Zynq-7000 SoC with an ARM 2x Cortex A9 CPU and a Kintex-7 FPGA. The form factor is the expandable PCIe-104. Peripherals: 1GB DDR3 memory, Quad SPI Flash, Gbit Ethernet, Micro SDHC, USB 2.0, dual role HDMI, PCIe 1x.",
        description: "12 layer PCB design for High Performance ...",
    },
    {
        _id: "100001",
        image: "https://www.kicad.org/made-with-kicad/librem-5-ext-con-rigid-flex-breakout-board/Librem_5_EXT_CON_Rigid_Flex_Breakout_Board_v0.6.png",
        projectName: "Purism: Librem 5 EXT CON",
        // description: "Small board that connects to the FPC connector J10 on the Librem 5 while the rigid section mounts inside of the phone using the center M.2 cover-plate threaded insert. The breakout board provides UART, I2C, SPI, two GPIOs, 3.3V, and 1.8V; making it really useful for custom projects",
        description: "Small board that connects to the FPC ...",
    },
    {
        _id: "100001",
        image: "https://www.kicad.org/made-with-kicad/nuco-v/nuco-v.jpg",
        projectName: "Dmitry: NUCO-V",
        // description: "NUCLEO-64© compatible development board for the STM32F7 and STM32H7 series with integrated Black Magic Probe",
        description: "NUCLEO-64© compatible development ...",
    },
    {
        _id: "100001",
        image: "https://www.kicad.org/made-with-kicad/pov_display_using_esp32/3d-view-ofpcb.jpg",
        projectName: "Jobit: POV Display",
        // description: "128 pixels resoultion and a maximum frame rate of 20FPS. This display is capable of displaying images as well as animation. The display is built around the ESP32 SoC, and used 74HC595 shift registers for controlling each pixel",
        description: "128 pixels resoultion and a maximum frame ...",
    },
    {
        _id: "100001",
        image: "https://www.kicad.org/made-with-kicad/spacedos/spacedos.jpg",
        projectName: "UST: SPACEDOS dosimeters",
        // description: "semiconductor detector based ionizing radiation spectrometers and dosimeters",
        description: "semiconductor detector based ionizing ...",
    },
    {
        _id: "100001",
        image: "https://www.kicad.org/made-with-kicad/upsat/upsat.jpg",
        projectName: "LSF: UPSat",
        // description: "2U Cubesat satellite constructed and delivered by Libre Space Foundation, started by University of Patras as part of the upsatQB50 mission with ID GR-02.",
        description: "2U Cubesat satellite constructed and ...",
    },
]

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

function ContentFull() {
    return (
        <div className="flex flex-1 items-center justify-center rounded-lg border shadow-sm" x-chunk="dashboard-02-chunk-1">
            <div className="container my-6s">
                <div className="grid md:grid-cols-2 sm:grid-cols-1 lg:grid-cols-4 gap-4">
                    {projects.map((project, index) => (
                        <Project
                            key={index}
                            _id={project._id}
                            image={project.image}
                            projectName={project.projectName}
                            description={project.description}
                        />
                    ))}
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
    // TODO: Control this variable from API requests 
    var hasProjects = true
    // hasProjects = false

    if(hasProjects) {
        return <ContentFull />
    } else {
        return <ContentEmpty />
    }
}