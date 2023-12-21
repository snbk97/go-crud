import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { MainLayout } from "./layout/main";
import { LoginPage } from "./pages/login";


function App() {

    //use react-router-dom to route to different pages

    const router = createBrowserRouter([
      {
        path: "/",
        element: <MainLayout/>,
      },
      {
        path: "/auth",
        element: <LoginPage/>,
      },
    ]);
  return (
    <div className="App">
      <RouterProvider router={router} />
    </div>
  );
}

export default App;
