import Layout from "../Layout.jsx";
import Setting from "../components/Setting.jsx";

const Settings = () => {
  return (
    <Layout>
      <div className={"flex flex-col max-w-7xl max-lg:max-w-3xl mx-auto"}>
        <h1>Settings</h1>
        <div className={"grid grid-cols-2 max-lg:grid-cols-1 gap-4"}>
          <div className="flex flex-col gap-y-5 my-4 py-4 bg-black bg-opacity-15 rounded-lg shadow-md">
            <h2 className={"text-2xl"}>Network Settings</h2>
            <Setting
              name="Request Private Connections"
              func={() => {}} // Add function to handle setting change
              description="When enabled, only accept connection requests from approved users."
            />
            <Setting
              name={"Hide Public Device Name"}
              func={() => {}}
              description={
                "Hide your device's name from being visible to others on the network."
              }
            />
            <Setting
              name={"Enable Automatic Resource Sharing"}
              func={() => {}}
              description={
                "Automatically share resources with trusted devices on the network."
              }
            />
            <Setting
              name={"Enable Automatic Resource Archiving"}
              func={() => {}}
              description={
                "Automatically archive resources after sharing them with other devices."
              }
            />
          </div>
          <div className="flex flex-col gap-y-5 my-4 py-4 bg-black bg-opacity-15 rounded-lg shadow-md">
            <h2 className={"text-2xl"}>Cache Settings</h2>
            <Setting
              name={"Enable Cache Compression"}
              func={() => {}}
              description={
                "Compress cached resources to save storage space on your device."
              }
            />
            <Setting
              name={"Enable Cache Encryption"}
              func={() => {}}
              description={
                "Encrypt cached resources to ensure data privacy and security."
              }
            />
            <Setting
              name={"Enable Cache Expiration"}
              func={() => {}}
              description={
                "Set expiration time for cached resources to manage storage usage."
              }
            />
            <Setting
              name={"Enable Cache Size Limit"}
              func={() => {}}
              description={
                "Set a maximum size limit for the cache to prevent storage overflow."
              }
            />
          </div>
          <div className="flex flex-col gap-y-5 my-4 py-4 bg-black bg-opacity-15 rounded-lg shadow-md">
            <h2 className={"text-2xl"}>Security Settings</h2>
            <Setting
              name={"Enable Firewall Protection"}
              func={() => {}}
              description={
                "Enable firewall protection to block unauthorized access to your device."
              }
            />
            <Setting
              name={"Enable Intrusion Detection"}
              func={() => {}}
              description={
                "Detect and prevent intrusion attempts on your device from malicious sources."
              }
            />
            <Setting
              name={"Enable Malware Scanning"}
              func={() => {}}
              description={
                "Scan and remove malware from cached resources to prevent infection."
              }
            />
            <Setting
              name={"Enable Data Encryption"}
              func={() => {}}
              description={
                "Encrypt data transmission to ensure secure communication with peers."
              }
            />
          </div>
          <div className="flex flex-col gap-y-5 my-4 py-4 bg-black bg-opacity-15 rounded-lg shadow-md">
            <h2 className={"text-2xl"}>Appearance Settings</h2>
            <Setting
              name={"Enable Dark Mode"}
              func={() => {}}
              description={
                "Switch to dark mode for a more comfortable viewing experience."
              }
            />
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default Settings;
