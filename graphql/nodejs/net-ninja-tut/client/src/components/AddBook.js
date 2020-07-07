import React, { useState } from "react";
import { useQuery, useMutation } from "@apollo/react-hooks";
import {
  getAuthorsQuery,
  addBookMutation,
  getBooksQuery,
} from "../queries/queries";

export default function AddBooks() {
  const { loading, error, data } = useQuery(getAuthorsQuery);
  const [newBook, setNewBook] = useState({
    name: "",
    genre: "",
    authorId: "",
  });
  const [addBookMut, { dataMutation }] = useMutation(addBookMutation);

  const displayAuthors = () => {
    if (loading) return <option disabled>Loading...</option>;
    if (error) return <option disabled>Error loading authors</option>;
    if (data) {
      return data.authors.map((author, index) => {
        return (
          <option key={index} value={author.id}>
            {" "}
            {author.name}{" "}
          </option>
        );
      });
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    addBookMut({
      variables: {
        name: newBook.name,
        genre: newBook.genre,
        authorId: newBook.authorId,
      },
      refetchQueries: [{ query: getBooksQuery }],
    });
  };

  return (
    <div>
      <form id="add-book" onSubmit={handleSubmit}>
        <div className="field">
          <label>Book Name</label>
          <input
            type="text"
            onChange={(e) => setNewBook({ ...newBook, name: e.target.value })}
            value={newBook.name}
          />
        </div>

        <div className="field">
          <label>Genre</label>
          <input
            type="text"
            onChange={(e) => setNewBook({ ...newBook, genre: e.target.value })}
            value={newBook.genre}
          />
        </div>

        <div className="field">
          <label>Author</label>
          <select
            onChange={(e) =>
              setNewBook({ ...newBook, authorId: e.target.value })
            }
          >
            <option>Select author</option>
            {displayAuthors()}
          </select>
        </div>

        <button>+</button>
      </form>
    </div>
  );
}
