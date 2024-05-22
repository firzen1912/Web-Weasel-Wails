import { useState, useEffect } from "react";
import Layout from "../Layout.jsx";

const Caching = () => {
  // State for storing URL list
  const [urlList, setUrlList] = useState([]);
  // State for input value
  const [urlInput, setUrlInput] = useState("");
  // State for output message
  const [outputMessage, setOutputMessage] = useState("");

  // Retrieve stored URL list from localStorage on component mount
  useEffect(() => {
    const storedUrlList = JSON.parse(localStorage.getItem("urlList"));
    if (storedUrlList) {
      setUrlList(storedUrlList);
    }
  }, []);

  // Function to update and store the URL list in localStorage
  const updateAndStoreUrlList = () => {
    localStorage.setItem("urlList", JSON.stringify(urlList));
  };

  // Function to clear a specific URL from the list
  const clearUrl = (url) => {
    const updatedUrlList = urlList.filter((u) => u !== url);
    setUrlList(updatedUrlList);
    updateAndStoreUrlList();
  };

  // Function to handle fetch button click
  const handleFetchButtonClick = () => {
    const urlRegex = /^(https?|http):\/\/[^\s$.?#].[^\s]*$/i;

    if (!urlInput || !urlRegex.test(urlInput)) {
      setOutputMessage("Please enter a valid URL.");
      return;
    }

    if (urlList.includes(urlInput)) {
      setOutputMessage("URL is already in the list.");
      return;
    }

    const updatedUrlList = [...urlList, urlInput];
    setUrlList(updatedUrlList);
    setUrlInput(""); // Clear input field

    updateAndStoreUrlList();

    const url =
      "http://localhost:8080/cache?path=" + encodeURIComponent(urlInput);

    fetch(url, { mode: "no-cors" })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.text();
      })
      .then((data) => {
        console.log("Cached URL:", data);
        setOutputMessage("Cached URL: " + data);
      })
      .catch((error) => {
        console.error("There was a problem with your fetch operation:", error);
      });
  };

  return (
    <Layout>
      <div
        id="downloadContainer"
        className="flex flex-col gap-y-5 max-w-3xl max-lg:max-w-lg mx-auto"
      >
        <h1 className={"text-3xl"}>Web Caching</h1>
        <div className={"input-container"}>
          <label htmlFor="urlInput">Enter web page URL:</label>
          <input
            type="text"
            id="urlInput"
            value={urlInput}
            onChange={(e) => setUrlInput(e.target.value)}
          />
          <button id="fetchButton" onClick={handleFetchButtonClick}>
            Download
          </button>
        </div>
        <p id="output" className={"response-message"}>
          {outputMessage}
        </p>
        <div className={"saved-pages"}>
          <h2 className={"text-2xl"}>Saved Webpages</h2>
          <p>
            Click on the URL to view the cached webpage. Click "Clear" to remove
            the URL from the list.
          </p>
          <ul id="urlList">
            {urlList.map((url) => (
              <li key={url}>
                <a
                  href={"http://localhost:8080/retrieve?path=" + url}
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  {url}
                </a>
                <button onClick={() => clearUrl(url)}>Clear</button>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </Layout>
  );
};

export default Caching;
