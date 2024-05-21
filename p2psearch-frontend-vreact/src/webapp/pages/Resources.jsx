import Layout from "../Layout.jsx";
import { FaGithub } from "react-icons/fa";
import { FaGolang } from "react-icons/fa6";

const Resources = () => {
  return (
    <Layout>
      <section className="content-section">
        <h1>Open Source Resources</h1>
        <div>
          <h2>Project Sources</h2>
          <p>
            Visit our GitHub repositories to access additional resources,
            documentation, and contribute to the project:
          </p>
          <div className={"flex justify-center"}>
            <ul className={"grid grid-cols-3 gap-x-10"}>
              <li>
                <a
                  className={"resource-item"}
                  href="https://github.com/alexkefer/p2pWebCaching"
                  target="_blank"
                >
                  <FaGithub className={"text-3xl"} />
                  <h4>Peer-to-Peer Networking Component</h4>
                </a>
              </li>
              <li>
                <a
                  className={"resource-item"}
                  href="https://github.com/alexkefer/webDownloader"
                  target="_blank"
                >
                  <FaGithub className={"text-3xl"} />
                  <h4>Web Caching Component</h4>
                </a>
              </li>
              <li>
                <a
                  className={"resource-item"}
                  href="https://pkg.go.dev/github.com/alexkefer/webDownloader"
                  target="_blank"
                >
                  <FaGolang className={"text-3xl"} />
                  <h4>Web Downloading Documentation</h4>
                </a>
              </li>
            </ul>
          </div>
        </div>
        <div>
          <h2>Additional Resources</h2>
          <p>Here are some documents that assisted us in research:</p>
        </div>
      </section>
    </Layout>
  );
};

export default Resources;
