import {HashRouter, Route, Routes} from "react-router-dom";
import Webapp from "./pages/Webapp.jsx";
import Settings from "./pages/Settings.jsx";
import Caching from "./pages/Caching.jsx";
import Tutorial from "./pages/Tutorial.jsx";
import Resources from "./pages/Resources.jsx";
import ErrorPage from "./pages/ErrorPage.jsx";

function App() {

    return (
        <div id="App">
            <HashRouter>
                <Routes>
                    <Route path="/" element={<Webapp />} />
                    <Route path="/settings" element={<Settings />} />
                    <Route path="/caching" element={<Caching />} />
                    <Route path="/tutorial" element={<Tutorial />} />
                    <Route path="/resources" element={<Resources />} />
                    <Route path="*" element={<ErrorPage />} />
                </Routes>
            </HashRouter>
        </div>
    )
}

export default App
