import DashboardContent from "@/components/content/dashboard-content";
import { AddProject } from "@/components/dashboard/add-project";

export function AllProjectsPage() {
    return (
        <>
            <div className="flex items-center justify-between">
                <h1 className="text-lg font-semibold md:text-2x1">Projects</h1>
				<AddProject />
            </div>
            <DashboardContent />
        </>
    )
}

export default AllProjectsPage