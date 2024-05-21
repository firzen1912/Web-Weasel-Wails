import React from "react";
import ReactDOM from "react-dom/client";
import { MemoryRouter as Router, Routes, Route } from "react-router-dom";
import ErrorPage from "./pages/ErrorPage.jsx";
import Webapp from "./pages/Webapp.jsx";
import Settings from "./pages/Settings.jsx";
import Caching from "./pages/Caching.jsx";
import Tutorial from "./pages/Tutorial.jsx";
import Resources from "./pages/Resources.jsx";
import "./webapp.css";

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <Router>
      <Routes>
        <Route path="/" element={<Webapp />} />
        <Route path="/settings" element={<Settings />} />
        <Route path="/caching" element={<Caching />} />
        <Route path="/tutorial" element={<Tutorial />} />
        <Route path="/resources" element={<Resources />} />
        <Route path="*" element={<ErrorPage />} />
      </Routes>
    </Router>
  </React.StrictMode>,
);
