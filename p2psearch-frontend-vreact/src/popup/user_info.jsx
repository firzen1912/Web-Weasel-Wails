import React, { useEffect, useState } from "react";
import "./popup.css";

function User_Info() {
  const [data, setData] = useState({
    deviceName: "",
    ipAddress: "",
    nearestNode: "",
  });

  useEffect(() => {
    // Function to generate random name
    const generateRandomName = () => {
      const names = ["MyNetwork1"];
      return names[Math.floor(Math.random() * names.length)];
    };

    // Function to generate random IP
    const generateRandomIP = () => {
      const baseIP = "192.168.0.";
      const randomOctet = Math.floor(Math.random() * 255) + 1; // Generate a random number between 1 and 255
      return baseIP + randomOctet;
    };

    // Function to generate random node
    const generateRandomNode = () => {
      const nodes = [""];
      return nodes[Math.floor(Math.random() * nodes.length)];
    };

    // Generate initial random data
    setData({
      deviceName: generateRandomName(),
      ipAddress: generateRandomIP(),
      nearestNode: generateRandomNode(),
    });
  }, []);

  return (
    <div>
      <label className="user-info">
        <span className="device-info-text">
          Public Device Name: {data.deviceName}
        </span>
        <span className="ip-adr-text">Node IP Address: {data.ipAddress}</span>
        <span className="neighbor-ip-text">
          Nearest Connection Node: {data.nearestNode}
        </span>
      </label>
    </div>
  );
}

export default User_Info;
