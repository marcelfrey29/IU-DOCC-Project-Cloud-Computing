import { Route, Routes } from "react-router-dom";

import AboutPage from "@/pages/about";
import IndexPage from "@/pages/index";

function App() {
    return (
        <Routes>
            <Route element={<IndexPage />} path="/" />
            <Route element={<AboutPage />} path="/about" />
        </Routes>
    );
}

export default App;
