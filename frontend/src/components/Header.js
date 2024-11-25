import React from "react";
import { NavLink } from "react-router";

export default function Header(props) {
  console.log(props);
  return (
    <nav style={styles.navContainer}>
      {/* NavLink makes it easy to show active states */}
      <NavLink
        style={styles.navLink}
        to="/"
        className={({ isActive }) => (isActive ? "active" : "")}
      >
        Home
      </NavLink>
      <NavLink style={styles.navLink} to="/login">
        Login
      </NavLink>
      <NavLink style={styles.navLink} to="/register">
        Signup
      </NavLink>
      <NavLink style={styles.navLink} to="/profile">
        Profile
      </NavLink>
    </nav>
  );
}

const styles = {
  navContainer: {
    padding: "10px",
    backgroundColor: "DodgerBlue",
  },
  navLink: {
    padding: "0 10px",
    textDecoration: "none",
  },
};
