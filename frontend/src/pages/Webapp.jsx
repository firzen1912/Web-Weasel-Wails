//import { useState } from 'react'
import Layout from "../Layout.jsx";
import { Link } from "react-router-dom";
import { FaExclamation } from "react-icons/fa";

const Webapp = () => {
  return (
    <Layout>
      <div className="flex flex-col gap-y-5 max-w-3xl max-lg:max-w-lg mx-auto">
        <h1 className={"text-2xl"}>Welcome to Peer-to-Peer Web Cache</h1>
        <p>
          This web application allows you to create a decentralized web cache
          network with your peers.
        </p>
        <div className={"flex flex-col gap-2 justify-between"}>
          <ul className={"styled-list"}>
            <h2 className={"text-xl"}>Features:</h2>
            <li>Decentralized caching of web content</li>
            <li>
              Reduced bandwidth usage by sharing cached content with peers
            </li>
            <li>Improved access speed to frequently visited websites</li>
            <li>User-friendly interface for managing cache settings</li>
          </ul>
        </div>
        <div className={"flex flex-col gap-2 justify-between"}>
          <ol className={"styled-list"}>
            <h2 className={"text-xl"}>How to Use:</h2>
            <li>Download and install the application on your device.</li>
            <li>Create or join a peer network.</li>
            <li>Start caching and sharing web content with your peers.</li>
            <div
              className={
                "flex flex-row gap-2 p-2.5 bg-black bg-opacity-20 shadow-md rounded-md"
              }
            >
              <FaExclamation className={"text-red-600 text-lg ml-1"} />
              <p>
                For more detailed instructions on how to use the application,
                please refer to the{" "}
                <Link
                  to="/tutorial"
                  className={"underline text-blue-700 hover:text-blue-800"}
                >
                  tutorial.
                </Link>
              </p>
            </div>
          </ol>
        </div>
        <p className={"flex justify-center mx-10"}>
          Get started now and enjoy seamless connectivity with your peers
          sharing capabilities with those you trust!
        </p>
        <div className={"changelog-section"}>
          <h3 className={"text-xl"}>Changelog:</h3>
          <ul className={"changelog-list"}>
            <li className={"changelog-list-item"}>
              <b>Version 1.0.0</b>
              <ul>
                <li>Initial release of the web application.</li>
                <li>
                  Basic functionality for creating and managing peer networks.
                </li>
              </ul>
            </li>
            <li className={"changelog-list-item"}>
              <b>Version 1.1.0</b>
              <ul>
                <li>
                  Improved user interface on the extension by switching to Vite
                  & React.
                </li>
              </ul>
            </li>
          </ul>
        </div>
      </div>
    </Layout>
  );
};

export default Webapp;
