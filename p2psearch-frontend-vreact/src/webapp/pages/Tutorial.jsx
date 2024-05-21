import Layout from "../Layout.jsx";
import { FaLongArrowAltRight } from "react-icons/fa";

const Tutorial = () => {
  return (
    <Layout>
      <div className="flex flex-col gap-y-5 max-w-3xl max-lg:max-w-lg mx-auto">
        <h2>Tutorial</h2>
        <p>
          This web application allows you to create a decentralized web cache
          network with your peers.
        </p>
        <div
          className={
            "list-inside p-2 rounded-md shadow-md bg-black bg-opacity-15"
          }
        >
          <h3 className={"mb-4"}>How to Install the extension:</h3>
          <ol className={"tutorial-ordered-list"}>
            <li>Clone the github repo to your local device.</li>
            <li>Open your chromium based browser and go to extensions.</li>
            <li>Turn on developer mode.</li>
            <img
              src="/images/tutorialImage1.png"
              alt="tutorial1"
              className={"rounded-xl p-2 mb-2 bg-black bg-opacity-15 mx-auto"}
            />
            <li>Click load unpacked.</li>
            <li>
              Navigate to
              <p className={"text-[16px] bg-black bg-opacity-10 rounded-md"}>
                ./p2pWebCaching/p2psearch-frontend
              </p>
              {/* Hopefully we can simplify this going forward */}
            </li>
            <li>Click "Select Folder".</li>
            <img
              src="/images/tutorialImage2.png"
              alt="tutorial2"
              className={"rounded-xl p-2 mb-2 bg-black bg-opacity-15 mx-auto"}
            />
          </ol>
          <p>
            You now should be able to open the extension and the website in your
            browser
          </p>
        </div>
        <div
          className={
            "list-inside p-2 rounded-md shadow-md bg-black bg-opacity-15"
          }
        >
          <h3>How to Connect Using the Extension:</h3>
          <ol className={"tutorial-ordered-list"}>
            <li>Open up your console.</li>
            <li>
              Navigate to
              <p>~/p2pWebCaching/p2psearch-backend</p>
            </li>
            <li>
              Enter the command:
              <ol
                className={"mt-2 px-2"}
                style={{
                  listStyleType: "lower-alpha",
                  listStylePosition: "inside",
                }}
              >
                <li>
                  If you are <u>hosting</u> the network:
                  <code
                    className={
                      "justify-center flex bg-black bg-opacity-10 rounded-md"
                    }
                  >
                    go run .
                  </code>
                </li>
                <li>
                  If you are <u>connecting</u> to an already hosted instance:
                  <code
                    className={
                      "justify-center flex bg-black bg-opacity-10 rounded-md"
                    }
                  >
                    go run . ip &lt;addr&gt;:&lt;port&gt;
                  </code>
                </li>
              </ol>
            </li>
            <img
              src="/images/tutorialImage3.png"
              alt="tutorial2"
              className={"rounded-xl p-2 mb-2 bg-black bg-opacity-15 mx-auto"}
            />
            <p className={"px-10"}>
              To find the port and ip, ask for them from the host you want to
              connect to. They should be in lines below the &ldquo;
              <code>go run .</code>&rdquo; command on the host's console. These
              will be marked with the{" "}
              <code className={"text-blue-600"}>[INFO]</code> tag.
            </p>
          </ol>
        </div>
        <div
          className={
            "list-inside p-2 rounded-md shadow-md bg-black bg-opacity-15 mb-10"
          }
        >
          <h3>How to Use Web Cache:</h3>
          <ol className={"tutorial-ordered-list"}>
            <li>
              In the console navigate to
              <p>~/p2pWebCaching/p2psearch-backend</p>
            </li>
            <li>
              Run the server by entering:
              <code
                className={
                  "justify-center flex bg-black bg-opacity-10 rounded-md"
                }
              >
                go run .
              </code>
            </li>
            <li>
              Choose a website url you wish to access through the p2p extension.
            </li>
            <li>Choose a website URL you wish to access through the cache.</li>
            <li>Navigate to the Caching page on the extension website.</li>
            <li>Click fetch.</li>
          </ol>
          <div
            className={"flex justify-between bg-black bg-opacity-10 rounded-md"}
          >
            <img
              src="/images/tutorialImage4.png"
              alt="tutorial4"
              className={"rounded-xl p-2 mb-2 h-48"}
            />
            <FaLongArrowAltRight className={"my-auto text-6xl"} />
            <img
              src="/images/tutorialImage5.png"
              alt="tutorial5"
              className={"rounded-xl p-2 mb-2 h-48"}
            />
          </div>
          <p>
            The URL should now be a clickable link for all peers and viewable.
            Once clicked, the URL on that page when visiting through the
            extension should be in similar format to "http://localhost:8080"
          </p>
        </div>
      </div>
    </Layout>
  );
};

export default Tutorial;
