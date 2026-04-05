"use client"

import { useEffect, useRef, useState } from "react"
import { DayPicker } from "react-day-picker"
import { format } from "date-fns"
import { ja } from "date-fns/locale"
import "react-day-picker/style.css"

export const DatePicker = () => {
  const [selected, setSelected] = useState<Date>(new Date())
  const [isOpen, setIsOpen] = useState(false)
  const containerRef = useRef<HTMLDivElement>(null)

  const formatted = format(selected, "yyyy年M月d日")
  const isoValue = format(selected, "yyyy-MM-dd")

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
      <label className="text-sm font-medium">日付</label>
      <div ref={containerRef} className="relative">
        <input type="hidden" name="occured_at" value={isoValue} />

        <button
          type="button"
          onClick={() => setIsOpen((prev) => !prev)}
          className="w-full flex items-center justify-between border rounded-lg px-3 py-2 bg-gray-50 text-sm focus:outline-none focus:ring-2 focus:ring-blue-400"
        >
          <span>{formatted}</span>
          <span className="text-gray-400">📅</span>
        </button>

        {isOpen && (
          <div className="absolute z-20 mt-1 bg-white border rounded shadow-xl animate-dropdown">
            <DayPicker
              mode="single"
              selected={selected}
              onSelect={(date) => {
                if (date) {
                  setSelected(date)
                  setIsOpen(false)
                }
              }}
              locale={ja}
              captionLayout="dropdown"
            />
          </div>
        )}
      </div>
    </div>
  )
}
