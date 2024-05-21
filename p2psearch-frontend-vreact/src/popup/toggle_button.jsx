import React, { useState } from "react";

function ToggleButton() {
  const [iconSrc, setIconSrc] = useState("../images/on_power_icon.png");
  const [deviceInfo, setDeviceInfo] = useState({
    deviceName: generateRandomName(),
    ipAddress: generateRandomIP(),
    nearestNode: generateRandomNode(),
  });

  const toggleIcon = () => {
    const newIconSrc = iconSrc.includes("on_power_icon.png")
      ? "../images/off_power_icon.png"
      : "../images/on_power_icon.png";
    setIconSrc(newIconSrc);

    const randomData = {
      deviceName: generateRandomName(),
      ipAddress: generateRandomIP(),
      nearestNode: generateRandomNode(),
    };

    toggleDeviceInfoVisibility(
      newIconSrc.includes("on_power_icon.png"),
      randomData,
    );
  };

  const toggleDeviceInfoVisibility = (isIconOn, data) => {
    const publicDeviceNameSpan = document.querySelector(".device-info-text");
    const ipAddressSpan = document.querySelector(".ip-adr-text");
    const nearestNodeSpan = document.querySelector(".neighbor-ip-text");

    if (isIconOn && data) {
      setDeviceInfo(data);
      publicDeviceNameSpan.textContent =
        "Public Device Name: " + data.deviceName;
      ipAddressSpan.textContent = "Node IP Address: " + data.ipAddress;
      nearestNodeSpan.textContent =
        "Nearest Connection Node: " + data.nearestNode;
    } else {
      setDeviceInfo({
        deviceName: "",
        ipAddress: "",
        nearestNode: "",
      });
      publicDeviceNameSpan.textContent = "Public Device Name: ";
      ipAddressSpan.textContent = "Node IP Address: ";
      nearestNodeSpan.textContent = "Nearest Connection Node: ";
    }
  };

  return (
    <div>
      <label className="icon">
        <div className="icon-container">
          <button id="iconButton" className="iconButton" onClick={toggleIcon}>
            <img className="icon-img" src={iconSrc} alt="Icon" />
          </button>
        </div>
      </label>
    </div>
  );
}

function generateRandomName() {
  const names = ["MyNetwork1"];
  return names[Math.floor(Math.random() * names.length)];
}

function generateRandomIP() {
  const baseIP = "192.168.0.";
  const randomOctet = Math.floor(Math.random() * 255) + 1;
  return baseIP + randomOctet;
}

function generateRandomNode() {
  const nodes = [""];
  return nodes[Math.floor(Math.random() * nodes.length)];
}

export default ToggleButton;
