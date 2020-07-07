import React, { useState } from "react";
import { useQuery } from "@apollo/react-hooks";
import { getBooksQuery } from "../queries/queries";
import BookDetails from "../components/BookDetails";

export default function BookList() {
  const { loading, data } = useQuery(getBooksQuery);
  const [selectedBookId, setSelectedBookId] = useState(null);

  const displayBooks = () => {
    return loading ? (
      <li>Loading..</li>
    ) : (
      data.books.map((book) => (
        <li
          key={book.id}
          onClick={(e) => {
            console.log(book.id);
            setSelectedBookId(book.id);
          }}
        >
          {book.name}
        </li>
      ))
    );
  };

  return (
    <div>
      <ul id="book-list">{displayBooks()}</ul>
      <BookDetails bookId={selectedBookId} />
    </div>
  );
}
