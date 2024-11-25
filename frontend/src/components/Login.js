import React, { useState } from "react";
import { useNavigate } from "react-router";

export default function Login() {
  const navigate = useNavigate();
  const [account, setAccount] = useState({ email: "", password: "" });
  const handleLogin = (event) => {
    event.preventDefault();
    fetch("http://localhost:4000/api/login", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: account.email,
        password: account.password,
      }),
    })
      .then((res) => res.json())
      .then((json) => {
        console.log(json["token"]);
        localStorage.setItem("token", json["token"]);
        navigate("/");
      })
      .catch((err) => {
        localStorage.removeItem("token");
        console.log(err);
      });
  };

  const handleAccount = (property, event) => {
    const accountCopy = { ...account };
    accountCopy[property] = event.target.value;
    setAccount(accountCopy);
  };
  return (
    <div>
      <form>
        <fieldset>
          <p>
            <label htmlFor="email">
              Email:
              <input
                onChange={(event) => handleAccount("email", event)}
                required
                id="email"
                name="email"
                autoFocus
              />
            </label>
          </p>
          <p>
            <label htmlFor="password">
              Password:
              <input
                onChange={(event) => handleAccount("password", event)}
                required
                name="password"
                type="password"
                id="password"
                autoComplete="current-password"
              />
            </label>
          </p>
          <p>
            <input type="submit" onClick={handleLogin} value="Sign In" />
          </p>
        </fieldset>
      </form>
    </div>
  );
}

export const styles = {
  Container: {
    justifyContent: "center",
  },
  avatar: {
    justifyContent: "center",
  },
};
