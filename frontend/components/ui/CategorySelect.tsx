"use client"

import { useEffect, useRef, useState } from "react"
import { CategoryResponse as Category } from "@/types/api"

type Props = {
  categories: Category[]
}

export const CategorySelect = ({ categories }: Props) => {
  const [isOpen, setIsOpen] = useState(false)
  const [selectedId, setSelectedId] = useState<string>(categories[0]?.id ?? "")
  const [focusedIndex, setFocusedIndex] = useState(0)
  const containerRef = useRef<HTMLDivElement>(null)

  const selectedName =
    categories.find((c) => c.id === selectedId)?.name ?? "選択してください"

  const handleSelect = (id: string, index: number) => {
    setSelectedId(id)
    setFocusedIndex(index)
    setIsOpen(false)
  }

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (!isOpen) {
      if (e.key === "Enter" || e.key === " ") {
        e.preventDefault()
        setIsOpen(true)
      }
      return
    }

    switch (e.key) {
      case "ArrowDown":
        e.preventDefault()
        setFocusedIndex((prev) => (prev + 1) % categories.length)
        break
      case "ArrowUp":
        e.preventDefault()
        setFocusedIndex((prev) =>
          prev === 0 ? categories.length - 1 : prev - 1
        )
        break
      case "Enter":
      case " ":
        e.preventDefault()
        handleSelect(categories[focusedIndex].id, focusedIndex)
        break
      case "Escape":
        setIsOpen(false)
        break
    }
  }

  // 外側クリックで閉じる
  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      if (
        containerRef.current &&
        !containerRef.current.contains(e.target as Node)
      ) {
        setIsOpen(false)
      }
    }
    document.addEventListener("mousedown", handleClickOutside)
    return () => document.removeEventListener("mousedown", handleClickOutside)
  }, [])

  return (
    <div className="flex flex-col gap-1">
      <label className="text-sm font-medium">カテゴリ</label>
      <div ref={containerRef} className="relative">
        <input type="hidden" name="category_id" value={selectedId} />

        <button
          type="button"
          onClick={() => setIsOpen((prev) => !prev)}
          onKeyDown={handleKeyDown}
          aria-haspopup="listbox"
          aria-expanded={isOpen}
          className="w-full flex items-center justify-between border rounded-lg px-3 py-2 bg-gray-50 text-sm focus:outline-none focus:ring-2 focus:ring-blue-400"
        >
          <span>{selectedName}</span>
          <span
            className={`transition-transform duration-200 ${isOpen ? "rotate-180" : "rotate-0"}`}
          >
            ▾
          </span>
        </button>

        {isOpen && (
          <ul
            role="listbox"
            onKeyDown={handleKeyDown}
            className="animate-dropdown absolute z-10 mt-1 w-full bg-white border rounded shadow-lg overflow-hidden"
          >
            {categories.map((c, i) => (
              <li
                key={c.id}
                role="option"
                aria-selected={c.id === selectedId}
                onMouseEnter={() => setFocusedIndex(i)}
                onClick={() => handleSelect(c.id, i)}
                className={`px-3 py-2 text-sm cursor-pointer transition-colors duration-100 ${
                  i === focusedIndex
                    ? "bg-blue-50 text-blue-700"
                    : "hover:bg-gray-50"
                }`}
              >
                {c.name}
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  )
}
