import React, { useState } from "react";
import { useNavigate } from "react-router";

export default function Todo() {
  const navigate = useNavigate();
  const [newTodo, setNewTodo] = useState();
  const handleNewTodo = (property, event) => {
    const newTodoCopy = { ...newTodo };
    newTodoCopy[property] = event.target.value;
    setNewTodo(newTodoCopy);
  };
  const handleSubmit = (event) => {
    event.preventDefault();
    fetch("http://localhost:5000/api/todos", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
        "x-access-token": localStorage.getItem("token"),
      },
      body: JSON.stringify({
        name: newTodo.name,
        description: newTodo.description,
      }),
    })
      .then((res) => res.json())
      .then((json) => {
        console.log(json);
        navigate("/");
      })
      .catch((err) => {
        console.log(err);
      });
  };
  return (
    <form>
      <fieldset>
        <legend>New Todo</legend>
        <p>
          <label htmlFor="name">
            Name:
            <input
              type="text"
              name="name"
              onChange={(event) => handleNewTodo("name", event)}
            />
          </label>
        </p>
        <p>
          <label htmlFor="description">
            Description:
            <input
              type="text"
              name="description"
              onChange={(event) => handleNewTodo("description", event)}
            />
          </label>
        </p>
        <p>
          <input type="submit" value="create" onClick={handleSubmit} />
        </p>
      </fieldset>
    </form>
  );
}
