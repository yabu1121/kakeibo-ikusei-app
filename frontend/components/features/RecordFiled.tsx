import { getCategories } from "@/lib/server/api"
import { FormLine } from "../ui/FormLine"
import { Button } from "../ui/Button"
import { CategorySelect } from "../ui/CategorySelect"
import { DatePicker } from "../ui/DatePicker"

export const RecordFiled = async () => {
  const categories = await getCategories()
  return (
    <div className="flex flex-col gap-4">
      <FormLine type="expense"/>
      <FormLine type="amount" inputProps={{min: "1"}}/>
      <DatePicker />
      <CategorySelect categories={categories} />
      <Button type="submit">記録する</Button>
    </div>
  )
}