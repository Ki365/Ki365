import { Button } from "@/components/ui/button";
import { PATH } from "@/routes/routes";
import { Link } from "react-router-dom";

import banner from "@/assets/banner.svg"

export function HomePage() {
    return (
        <>
            <div className="flex flex-col m-auto justify-center w-screen min-h-screen items-center space-y-4">
                <img src={banner} className="p-6"></img>
                <div>This is a temporary home page, press the button below to login.</div>
                <div>We will update it shortly...</div>
                <Link to={PATH.GENERAL.LOGIN}>
                    <Button>
                Go to login page
                    </Button>
                </Link>
            </div>
        </>
    )
}