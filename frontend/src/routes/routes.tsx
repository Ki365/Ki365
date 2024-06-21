import LoginPage from "@/pages/login";
import DashboardPage from "@/pages/dashboard/dashboard";
import ProjectPage from "@/pages/dashboard/projects/project";
import { AllProjectsPage } from "@/pages/dashboard/projects/all-projects";
import { HomePage } from "@/pages/homepage";
import AdminPage from "@/pages/dashboard/admin/admin";
import ComponentsPage from "@/pages/dashboard/components/components";
import MembersPage from "@/pages/dashboard/members/members";
import PreferencesPage from "@/pages/dashboard/preferences/preferences";

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
        COMPONENTS: "/dashboard/components/",
        MEMBERS: "/dashboard/members/",
        PREFERENCES: "/dashboard/preferences/",
        SURPRISE: "/dashboard/surprise/",
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
    { path: PATH.ADMIN.ADMIN_SETTINGS, element: <AdminPage /> },

    // USER ROUTES
    { path: PATH.USER.HOME, element: <DashboardPage /> },
    { path: PATH.USER.ALL_PROJECTS, element: <AllProjectsPage /> },
    { path: PATH.USER.PROJECT, element: <ProjectPage /> },
    { path: PATH.USER.COMPONENTS, element: <ComponentsPage /> },
    { path: PATH.USER.MEMBERS, element: <MembersPage /> },
    { path: PATH.USER.PREFERENCES, element: <PreferencesPage /> },
    { path: PATH.USER.SURPRISE, element: <>Surprise!! </> },
]
