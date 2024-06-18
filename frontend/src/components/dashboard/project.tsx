import { CircleUserIcon, FullscreenIcon } from "lucide-react"
import React from "react"
import { Button } from "@/components/ui/button"

// TODO: turn any into types for typescript to do its job
const Project = (props: any) => {
    return (

        <div className="w-full relative group border rounded-lg">
            <div className="max-w-80 max-h-80 relative overflow-y-hidden ">
                <a href="/dashboard/project" >
                    <div className="h-52 flex justify-center pt-2">
                        <img className="max-h-60 p-4" src={props.image}></img>
                    </div>
                </a>
                <div className="w-full h-16 absolute bg-white -bottom-[130px] group-hover:bottom-0 duration-700 border-t rounded-lg">
                    <ul className="h-full flex flex-col items-end justify-center gap-2 font-titleFont px-4">
                        <li className="text-[#767676] hover:text-primeColor text-sm font-normal border-b-[1px] border-b-gray-200 hover:border-b-primColor flex items-center justify-start gap-2 hover:cursor-pointer pb-1 duration-300 w-full">
                            <span className="p-2">
                                <FullscreenIcon size={16}>b</FullscreenIcon>
                            </span>
                            Expand quick view
                        </li>
                    </ul>
                </div>
            </div>
            <div className="max-w-80 py-6 flex flex-col gap-1 px-4">
                <div className="flex items-center justify-between font-medium">
                    <h2 className="text-md font-bold">{props.projectName}</h2>
                    <Button variant="secondary" size="icon" className="rounded-full h-8 w-8">
                        <a href="#" >
                            <CircleUserIcon size={16} />
                            <span className="sr-only">User profile</span>
                        </a>
                    </Button>
                    {/* <p className="text-[#767676 text-[14px]">{props.price}</p> */}
                </div>
                <div>
                    <p className="text-[#767676 text-[14px]">{props.description}</p>
                </div>
            </div>
        </div>
    )
}

export default Project