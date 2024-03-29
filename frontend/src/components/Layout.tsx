import { Outlet } from "react-router-dom";
import { ToastContainer } from "react-toastify";
import "react-toastify/ReactToastify.min.css";

import Header from "./Header";

const Layout = () => {
  return (
    <>
      <Header />
      <main>
        <Outlet />
      </main>
      <ToastContainer
        position="bottom-right"
        autoClose={5000}
        newestOnTop={false}
        draggable
        pauseOnHover={false}
        theme="colored"
      />
    </>
  );
};

export default Layout;
