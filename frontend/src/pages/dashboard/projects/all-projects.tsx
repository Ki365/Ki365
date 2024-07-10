import DashboardContent from "@/components/content/dashboard-content";
import { AddProject } from "@/components/dashboard/add-project";
import { createContext, useState } from "react";

export interface IType {
	user: boolean,
	setUser: (user: boolean) => void,
}

export const UserContext = createContext<IType | null>(null)

export function AllProjectsPage() {
	const [user, setUser] = useState(false)

    return (
        <>
			<UserContext.Provider value={{ user: user, setUser: setUser }} >
				<div className="flex items-center justify-between">
					<h1 className="text-lg font-semibold md:text-2x1">Projects</h1>
					<AddProject />
				</div>
				<DashboardContent />
			</UserContext.Provider>
        </>
    )
}

export default AllProjectsPage