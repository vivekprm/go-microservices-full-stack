import { Route, Routes } from "react-router";
import Login from "./components/Login";
import Profile from "./components/Profile";
import Home from "./components/Home";
import Header from "./components/Header";
import Register from "./components/Register";

function App() {
  return (
    <>
    <Header>
    </Header>
    <Routes>
        <Route index element={<Home />} />
        <Route path="login" element={<Login />} />
        <Route path="profile" element={<Profile />} />
        <Route path="register" element={<Register />} />
    </Routes>
    </>
  );
}

export default App;
