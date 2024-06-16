import { createBrowserRouter, RouterProvider, Outlet} from "react-router-dom"

import ROUTES from '@/routes/routes'

function App() {
  const router = createBrowserRouter([
    {
      element: <Outlet />,
      children: ROUTES
    }
  ])

  return (
    <>
      <RouterProvider router={router} />
    </>
  )
}

export default App
