"use client";

import { User } from "@/types/interfaces";
import { ColumnDef } from "@tanstack/react-table";
// import { Eye } from "lucide-react";

export const columns: ColumnDef<User>[] = [
  {
    accessorFn: (row) => row.fullname,
    id: "fullname",
    header: "Full Name",
  },
  {
    accessorFn: (row) => row.email,
    id: "email",
    header: "Email Address",
  },
  {
    accessorFn: (row) => row.badge_type,
    id: "badge_type",
    header: "Badge Type",
  },
  {
    accessorFn: (row) => row.status,
    id: "status",
    header: "Status",
    cell: ({ row }) => {
      const status = row.original.status;
      return <span className={`${status === "pending" ? "text-celebut-gold" : status === "approved" ? "text-celebut-green" : "text-celebut-red"} capitalize`}>{status}</span>;
    },
  },
  {
    accessorFn: (row) => `${row.state}, ${row.country}`,
    id: "location",
    header: "Location",
    cell: ({ row }) => {
      const state = row.original.state as string;
      const country = row.original.country as string;

      return (
        <div>
          <p>{state}</p>
          <p className="text-[#999999]">{country}</p>
        </div>
      );
    },
  },
  {
    accessorFn: (row) => row.document_type,
    id: "document_type",
    header: "Document Type",
  },
  {
    accessorFn: (row) => row.id,
    id: "id",
    header: "View",
  },
];
