import { Button } from "@/components/ui/button";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuSeparator, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet";
import { Separator } from "@/components/ui/separator";
import { BellOffIcon, BookOpenIcon, CircleUserIcon, CpuIcon, FoldersIcon, MenuIcon, PartyPopperIcon, SearchIcon, ShieldEllipsisIcon, SunMoonIcon, UserRoundCogIcon, UsersIcon } from "lucide-react";

import banner from '@/assets/banner.png'

import { SiGithub } from '@icons-pack/react-simple-icons'
import { Toaster } from "sonner";

import { Link, Outlet } from "react-router-dom";
import { PATH } from "@/routes/routes";

export default function DashboardTemplate() {
    return (
        <div className="grid min-h-screen w-screen md:grid-cols-[220px_1fr] lg:grid-cols-[280px_1fr]">
            <div className="hidden border-r bg-muted/40 md:block">
                <div className="flex h-full max-h-screen flex-col gap-2">
                    <div className="flex h-14 items-center border-b px-4 lg:h-[60px] lg:px-6 py-8">
                        <a href="/" className="flex items-center gap-2 font-semibold">
                            <img src={banner} className="h-10" />
                            {/* <img src={ico} className="h-6 w-6" />
                        <span className="">Ki365</span> */}
                        </a>
                        <Button variant="outline" size="icon" className="ml-auto h-8 w-8">
                            <BellOffIcon className="h-4 w-4" />
                            <span className="sr-only">Toggle Notifications</span>
                        </Button>
                    </div>
                    <div className="flex-1">
                        <nav className="grid items-start px-2 text-sm font-medium lg:px-4">
                            <br className="" />

                            <Link to={PATH.USER.HOME}
                                className="flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-muted-foreground">
                                <BookOpenIcon className="h-4 w-4" />
                                Home
                            </Link>
                            {/* <br className="" /> */}
                            <Separator className="my-2" />

                            <Link to={PATH.USER.ALL_PROJECTS}
                                className="flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-primary bg-muted">
                                <FoldersIcon className="h-4 w-4" />
                                Projects
                            </Link>

                            <a href="#"
                                className="flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-muted-foreground">
                                <CpuIcon className="h-4 w-4" />
                                Components
                            </a>
                            <Separator className="my-2" />

                            <a href="#"
                                className="flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-muted-foreground">
                                <UsersIcon className="h-4 w-4" />
                                Members
                            </a>
                            <a href="#"
                                className="flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-muted-foreground">
                                <UserRoundCogIcon className="h-4 w-4" />
                                Preferences
                            </a>
                            <Separator className="my-2" />
                            <a href="#"
                                className="flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-muted-foreground">
                                <ShieldEllipsisIcon className="h-4 w-4" />
                                Admin
                            </a>
                            <Separator className="my-2" />

                            <br className="" />
                            <br className="" />
                            <a href="#"
                                className="flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-muted-foreground">
                                <PartyPopperIcon className="h-4 w-4" />
                                Surprise Me!
                            </a>
                        </nav>
                    </div>
                </div>

            </div>
            <div className="flex flex-col">
                <header className="flex h-14 items-center gap-4 border-b bg-muted/40 px-4 lg:h-[60px] lg:px-6 py-8">
                    <Sheet>
                        <SheetTrigger asChild>
                            <Button variant="outline" size="icon" className="shrink-0 md:hidden">
                                <MenuIcon className="h-5 w-5" />
                                <span className="sr-only">Toggle navigation Menu</span>
                            </Button>
                        </SheetTrigger>
                        <SheetContent side="left" className="flex flex-col">
                            <div className="flex h-14 items-center border-b px-4 lg:h-[60px] lg:px-6">
                                <a href="/" className="flex items-center gap-2 font-semibold">
                                    <img src={banner} className="h-10" />
                                    {/* <img src={ico} className="h-6 w-6" />
                        <span className="">Ki365</span> */}
                                </a>
                                <Button variant="outline" size="icon" className="ml-auto h-8 w-8">
                                    <BellOffIcon className="h-4 w-4" />
                                    <span className="sr-only">Toggle Notifications</span>
                                </Button>
                            </div>
                            <nav className="grid items-start px-2 text-sm font-medium lg:px-4">
                                <Link to={PATH.USER.HOME}
                                    className="flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-muted-foreground">
                                    <BookOpenIcon className="h-4 w-4" />
                                    Home
                                </Link>
                                {/* <Separator className="my-2"/> */}
                                <Link to={PATH.USER.ALL_PROJECTS}
                                    className="flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-primary bg-muted">
                                    <FoldersIcon className="h-4 w-4" />
                                    Projects
                                </Link>
                                <a href="#"
                                    className="flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-muted-foreground">
                                    <CpuIcon className="h-4 w-4" />
                                    Components
                                </a>
                                <Separator className="my-2" />

                                <a href="#"
                                    className="flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-muted-foreground">
                                    <UsersIcon className="h-4 w-4" />
                                    Members
                                </a>
                                <a href="#"
                                    className="flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-muted-foreground">
                                    <UserRoundCogIcon className="h-4 w-4" />
                                    Preferences
                                </a>
                                <Separator className="my-2" />
                                <a href="#"
                                    className="flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-muted-foreground">
                                    <ShieldEllipsisIcon className="h-4 w-4" />
                                    Admin
                                </a>
                                <Separator className="my-2" />

                                <br className="" />
                                <br className="" />
                                <a href="#"
                                    className="flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-muted-foreground">
                                    <PartyPopperIcon className="h-4 w-4" />
                                    Surprise Me!
                                </a>
                            </nav>
                        </SheetContent>
                    </Sheet>
                    <div className="w-full flex-1">
                        <form>
                            <div className="relative">
                                <SearchIcon className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
                                <Input type="search" placeholder="Search site..." className="w-full appearance-none bg-background pl-10 shadow-none md:w-2/3 lg:w-1/2"></Input>
                            </div>
                        </form>
                    </div>
                    <Button variant="secondary" size="icon" className="rounded-full">
                        <a href="https://github.com/Ki365/Ki365/" >
                            <SiGithub title='Ki365 Github Link' className="h-5 w-5" />

                            <span className="sr-only">Toggle dark/light mode</span>
                        </a>
                    </Button>
                    <Button variant="secondary" size="icon" className="rounded-full" title='Dark/light mode' >
                        <SunMoonIcon className="h-5 w-5" />
                        <span className="sr-only">Toggle dark/light mode</span>
                    </Button>
                    <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                            <Button variant="secondary" size="icon" className="rounded-full" title='Profile'>
                                <CircleUserIcon className="h-6 w-6" />
                                <span className="sr-only">Toggle user menu</span>
                            </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent align="end">
                            <DropdownMenuItem>My Account</DropdownMenuItem>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem>Settings</DropdownMenuItem>
                            <DropdownMenuItem>Support</DropdownMenuItem>
                            <DropdownMenuSeparator />
                                <Link to={PATH.GENERAL.LOGIN}>
                            <DropdownMenuItem>
                                    Logout
                            </DropdownMenuItem>
                                </Link>
                        </DropdownMenuContent>
                    </DropdownMenu>
                </header>
                <main className="flex flex-1 flex-col gap-4 p-4 lg:gap-6 lg:p-6">

                    <Outlet />

                </main>
                <Toaster />
            </div >
        </div >
    )
}