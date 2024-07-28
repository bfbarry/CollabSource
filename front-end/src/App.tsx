import './App.css';
import './Pages/HomePage/HomePage';
import HomePage from './Pages/HomePage/HomePage';
import NavBar from './Pages/NavBar/NavBar';
import CreateProject from './Pages/CreateProject/CreateProject';
import ErrorPage from './error-page';
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import { SignedInContext } from './context/SignedInContext';
import { useState } from 'react';

const router = createBrowserRouter([
  {
    path: "/",
    element: <HomePage/>,
    errorElement: <ErrorPage/>
  },
  {
    path: "/explore",
    element:<div>Explore!</div>,
    errorElement: <ErrorPage/>
  },
  {
    path: "/about",
    element: <div>about</div>,
    errorElement: <ErrorPage/>
  },
  {
    path: "/project/:id",
    element: <div>Project page</div>,
    errorElement: <ErrorPage/>
  },
  {
    path: "/create_project",
    element: <CreateProject/>,
    errorElement: <ErrorPage/>
  },
]);

function App() {

  const [signedInUser, setSignedInUser] = useState((localStorage.getItem("access_token") == null ? false : true));

  return (
    <div>
      <SignedInContext.Provider value={{signedInUser, setSignedInUser}}>
        <NavBar/>
        <RouterProvider router={router}/>
      </SignedInContext.Provider>
    </div>
    
  );
}

export default App;
