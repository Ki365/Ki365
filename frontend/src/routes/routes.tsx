import LoginPage from "@/pages/login";
import DashboardPage from "@/pages/dashboard";

const PATH = {
    GENERAL: {
        HOME: "/",
        LOGIN: "/login",
    },

    ADMIN: {
        ADMIN_SETTINGS: "/admin/settings",
    },

    USER: {
        HOME: "/dashboard",
        PROJECT: "/dashboard/project",
    },

    INVITEE: {
        PROJECT: "/invitee/project",
    },
}

const ROUTES = [
    // GENERAL ROUTES
    { path: PATH.GENERAL.HOME, element: <div>home, go to ./login to login.</div> },
    { path: PATH.GENERAL.LOGIN, element: <LoginPage /> },

    // ADMIN ROUTES
    { path: PATH.ADMIN.ADMIN_SETTINGS, element: <DashboardPage /> },

    // USER ROUTES
    { path: PATH.USER.HOME, element: <DashboardPage /> },
    { path: PATH.USER.PROJECT, element: <DashboardPage /> },

    // INVITEE ROUTES
    { path: PATH.INVITEE.PROJECT, element: <DashboardPage /> },
]

export default ROUTES