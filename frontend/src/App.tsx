import { createBrowserRouter, RouterProvider, Outlet} from "react-router-dom"

import { ROUTES_GENERAL, ROUTES_DASHBOARD } from '@/routes/routes'
import DashboardTemplate from "./templates/dashboard-template"

function App() {
  const router = createBrowserRouter([
    {
      element: <Outlet />,
      errorElement: <div className="flex justify-center text-lg">404 not found</div>,
      children: [
        {
          element: <Outlet />,
          children: ROUTES_GENERAL,
        },
        {
          element: <DashboardTemplate />,
          children: ROUTES_DASHBOARD,
        },
      ]
    }
  ])

  return (
    <>
      <RouterProvider router={router} />
    </>
  )
}

export default App
