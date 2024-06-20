import LoginPage from "@/pages/login";
import DashboardPage from "@/pages/dashboard";
import ProjectPage from "@/pages/project";
import { AllProjectsPage } from "@/pages/all-projects";
import { HomePage } from "@/pages/homepage";

export const PATH = {
    GENERAL: {
        HOME: "/",
        LOGIN: "/login/",
    },

    ADMIN: {
        ADMIN_SETTINGS: "/admin/settings/",
    },

    USER: {
        HOME: "/dashboard/",
        ALL_PROJECTS: "/dashboard/projects/",
        PROJECT: "/dashboard/project/",
    },

    INVITEE: {
        PROJECT: "/invitee/project/",
    },
}

export const ROUTES_GENERAL = [
    // GENERAL ROUTES
    { path: PATH.GENERAL.HOME, element: <HomePage /> },
    { path: PATH.GENERAL.LOGIN, element: <LoginPage /> },

    // INVITEE ROUTES
    { path: PATH.INVITEE.PROJECT, element: <DashboardPage /> },
]

export const ROUTES_DASHBOARD = [
    // ADMIN ROUTES
    { path: PATH.ADMIN.ADMIN_SETTINGS, element: <DashboardPage /> },

    // USER ROUTES
    { path: PATH.USER.HOME, element: <DashboardPage /> },
    { path: PATH.USER.ALL_PROJECTS, element: <AllProjectsPage /> },
    { path: PATH.USER.PROJECT, element: <ProjectPage /> },
]
