import React, { useState, useEffect } from "react";
import "./popup.css";

function Popup() {
  const [deviceInfo, setDeviceInfo] = useState({
    deviceName: "",
    ipAddress: "",
    nearestNode: ""
  });
  const [peerAddress, setPeerAddress] = useState(""); // State to store peer address

  useEffect(() => {
    clearLocalStorage();
    fetchAndDisplayHostname();
    fetchAndDisplayNodeIPAddress();
    fetchAndDisplayPeersIPAddress();
  }, []);

  const clearLocalStorage = () => {
    localStorage.clear();
  };

  const fetchAndDisplayHostname = () => {
    fetch("http://localhost:8080/hostname")
      .then(response => {
        if (response.ok) {
          return response.text();
        } else {
          throw new Error("Failed to fetch hostname data");
        }
      })
      .then(hostname => {
        setDeviceInfo(prevState => ({
          ...prevState,
          deviceName: hostname
        }));
        localStorage.setItem("hostname", hostname);
      })
      .catch(error => {
        console.error("Error:", error.message);
      });
  };

  const fetchAndDisplayNodeIPAddress = () => {
    fetch("http://localhost:8080/peers")
      .then(response => {
        if (response.ok) {
          return response.text();
        } else {
          throw new Error("Failed to fetch node IP address data");
        }
      })
      .then(nodeIPAddress => {
        const lines = nodeIPAddress.split("\n");
        const firstLine = lines[0].trim();
        setDeviceInfo(prevState => ({
          ...prevState,
          ipAddress: firstLine
        }));
        localStorage.setItem("nodeIPAddress", firstLine);
      })
      .catch(error => {
        console.error("Error:", error.message);
      });
  };

  const fetchAndDisplayPeersIPAddress = () => {
    fetch("http://localhost:8080/peers")
      .then(response => {
        if (response.ok) {
          return response.text();
        } else {
          throw new Error("Failed to fetch peer IP address data");
        }
      })
      .then(nodeIPAddress => {
        const lines = nodeIPAddress.split("\n");
        const otherLines = lines.slice(1).map(line => line.trim());
        setDeviceInfo(prevState => ({
          ...prevState,
          nearestNode: otherLines.join(", ")
        }));
        localStorage.setItem("peerIPAddress", otherLines.join(", "));
      })
      .catch(error => {
        console.error("Error:", error.message);
      });
  };

  const toggleIcon = () => {
    // Assuming the icon is toggled using local state or deviceInfo
    // Adjust this according to your actual logic
  };

  const handlePeerAddressChange = (event) => {
    setPeerAddress(event.target.value);
  };

  const connectToPeer = () => {
    var connectURL = 'http://localhost:8080/connect?path=' + peerAddress;

    fetch(connectURL)
      .then(response => {
        if (response.ok) {
          console.log('Connected to peer successfully');
        } else {
          throw new Error('Failed to connect to peer');
        }
      })
      .catch(error => {
        console.error('Error:', error.message);
      });
  };

  return (
    <div>
      <label className="user-info">
        <span className="device-info-text">
          Public Device Name: {deviceInfo.deviceName}
        </span>
        <span className="ip-adr-text">
          Host IP Address: {deviceInfo.ipAddress}
        </span>
        <span className="neighbor-ip-text">
          Peers IP Address: {deviceInfo.nearestNode}
        </span>
      </label>

      <label className="icon">
        <div className="icon-container">
          <button id="iconButton" className="iconButton" onClick={toggleIcon}>
            <img
              className="icon-img"
              src="../../images/on_power_icon.png"
              alt="Icon"
            />
          </button>
        </div>
      </label>

      <div className="input-container">
        <input
          type="text"
          id="peerAddressInput"
          placeholder="Enter peer IP address"
          value={peerAddress}
          onChange={handlePeerAddressChange}
        />
        <button onClick={connectToPeer}>Connect</button>
      </div>

      <button
        onClick={() => window.open("webapp.html", "_blank")}
        type="button"
        className="custom-button"
      >
        Go to Web App
      </button>

      <label className="switch">
        <span className="label-text">Private Connection Request</span>
        <div className="slider-container">
          <input type="checkbox" id="toggleButton1" className="switch-input" />
          <span className="slider"></span>
        </div>
      </label>

      <label className="switch">
        <span className="label-text">Hide Public Device Name</span>
        <div className="slider-container">
          <input type="checkbox" id="toggleButton2" className="switch-input" />
          <span className="slider"></span>
        </div>
      </label>

      <label className="switch">
        <span className="label-text">Automatic Resource Sharing</span>
        <div className="slider-container">
          <input type="checkbox" id="toggleButton3" className="switch-input" />
          <span className="slider"></span>
        </div>
      </label>

      <label className="switch">
        <span className="label-text">Automatic Resource Archiving</span>
        <div className="slider-container">
          <input type="checkbox" id="toggleButton4" className="switch-input" />
          <span className="slider"></span>
        </div>
      </label>
    </div>
  );
}

export default Popup;
