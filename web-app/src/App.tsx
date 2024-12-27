import { Route, Routes } from "react-router-dom";

import AboutPage from "@/pages/about";
import IndexPage from "@/pages/index";
import TravelGuideDetailPage from "./pages/travel-guide";
import TravelGuidesListPage from "./pages/travel-guides-list";

function App() {
    return (
        <Routes>
            <Route element={<IndexPage />} path="/" />
            <Route element={<AboutPage />} path="/about" />
            <Route element={<TravelGuidesListPage />} path="/travel-guides" />
            <Route
                element={<TravelGuideDetailPage />}
                path="/travel-guides/:id"
            />
        </Routes>
    );
}

export default App;
