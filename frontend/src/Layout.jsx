import Navigation from "./components/Navigation";
import Footer from "./components/Footer";

//eslint-disable-next-line
const Layout = ({ children }) => {
  return (
    <div
      className={
        "min-h-screen max-w-screen bg-gradient-to-br from-blue-200 to-purple-200"
      }
    >
      <div className={"flex flex-row"}>
        <div>
          <Navigation />
        </div>
        <div className={"flex flex-col pt-4 w-full min-h-screen"}>
          <main className={"flex-grow mx-4"}>{children}</main>
          <div className="my-auto bg-black bg-opacity-30">
            <Footer />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Layout;
