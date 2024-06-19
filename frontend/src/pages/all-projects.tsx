import DashboardContent from "@/components/content/dashboard-content";

export function AllProjectsPage() {
    return (
        <>
            <div className="flex items-center">
                <h1 className="text-lg font-semibold md:text-2x1">Projects</h1>
            </div>
            <DashboardContent />
        </>
    )
}

export default AllProjectsPage