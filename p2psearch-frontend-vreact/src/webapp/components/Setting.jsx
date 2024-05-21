import Switch from "react-switch";
import { useState } from "react";

// eslint-disable-next-line react/prop-types
const Setting = ({ name, func, description }) => {
  const [toggleButton, setToggleButton] = useState(false);
  const handleChange = () => {
    setToggleButton(!toggleButton);
    func();
  };

  return (
    <div className="flex flex-col px-2 p-2">
      <div className="justify-between bg-black bg-opacity-15 rounded-md shadow-md mb-4">
        <div>
          <div className={"flex justify-between p-4"}>
            <p className={"mt-1"}>{name}</p>
            <Switch
              aria-label={name}
              checked={toggleButton}
              onChange={() => handleChange()}
              onColor="#a2e4b8"
              onHandleColor="#208E31"
              handleDiameter={30}
              uncheckedIcon={false}
              checkedIcon={false}
              boxShadow="0px 1px 5px rgba(0, 0, 0, 0.6)"
              activeBoxShadow="0px 0px 1px 2px rgba(0, 0, 0, 0.2)"
              height={20}
              width={48}
              id="setting-switch"
            />
          </div>
        </div>
      </div>
      <p>{description}</p>
    </div>
  );
};

export default Setting;
