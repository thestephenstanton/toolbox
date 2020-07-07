import React from "react";
import { useQuery } from "@apollo/react-hooks";
import { getBookQuery } from "../queries/queries";

export default function BookDetails(props) {
  const { loading, data } = useQuery(getBookQuery, {
    variables: { id: props.bookId },
  });

  const { book } = data ? data : { book: null };

  return book ? (
    <div id="book-details">
      <h2>{book.name}</h2>
      <p>{book.genre}</p>
      <p>{book.author.name}</p>
      <p>All books by this author</p>
      <ul className="other-books">
        {book.author.books.map((item) => {
          return <li key={item.id}>{item.name}</li>;
        })}
      </ul>
    </div>
  ) : (
    <div id="book-details">
      <p>no book selected</p>
    </div>
  );
}
