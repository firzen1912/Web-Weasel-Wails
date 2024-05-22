import { Link } from "react-router-dom";
import { Sidebar, Menu, MenuItem } from "react-pro-sidebar";
import { useEffect, useState } from "react";
import { FaHome, FaRegQuestionCircle } from "react-icons/fa";
import { FaGear } from "react-icons/fa6";
import { MdOutlineStorage } from "react-icons/md";
import { IoDocumentTextOutline, IoMenu } from "react-icons/io5";

const Navigation = () => {
  const [collapsed, setCollapsed] = useState(false);

  const collapseSidebar = () => {
    setCollapsed(!collapsed);
  };

  return (
    <div className={"h-full"}>
      <Sidebar
        width={"225px"}
        collapsedWidth={"60px"}
        collapsed={collapsed}
        rootStyles={{
          zIndex: 1000,
          height: "inherit",
          cursor: "pointer",
        }}
        backgroundColor={"rgba(0, 0, 0, 0.10)"}
        className={"bg-black bg-opacity-25"}
        onClick={collapseSidebar}
      >
        <Menu
          iconShape="square"
          menuItemStyles={{
            button: ({ active }) => {
              return {
                border: "none",
                padding: "0.75rem",
                width: "100%",
                ["&:hover"]: {
                  backgroundColor: "rgba(0, 0, 0, 0.15)",
                },
                ["&.active"]: {
                  backgroundColor: "rgba(0, 0, 0, 0.15)",
                },
              };
            },
          }}
        >
          <div className={"font-semibold text-xl flex justify-center my-2"}>
            {collapsed ? (
              <IoMenu className={"text-4xl"} />
            ) : (
              <div className={"flex gap-2"}>
                <h3>P2P Web Cache</h3>
              </div>
            )}
          </div>
          <MenuItem
            icon={<FaHome />}
            component={<Link to={"/"} className={"nav-item"} />}
          >
            Home
          </MenuItem>
          <MenuItem
            icon={<FaGear />}
            component={<Link to={"/settings"} className={"nav-item"} />}
          >
            Settings
          </MenuItem>
          <MenuItem
            icon={<MdOutlineStorage />}
            component={<Link to={"/caching"} className={"nav-item"} />}
          >
            Cache
          </MenuItem>
          <MenuItem
            icon={<FaRegQuestionCircle />}
            component={<Link to={"/tutorial"} className={"nav-item"} />}
          >
            Tutorial
          </MenuItem>
          <MenuItem
            icon={<IoDocumentTextOutline />}
            component={<Link to={"/Resources"} className={"nav-item"} />}
          >
            Resources
          </MenuItem>
        </Menu>
      </Sidebar>
    </div>
  );
};

export default Navigation;
