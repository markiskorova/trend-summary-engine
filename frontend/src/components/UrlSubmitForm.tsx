import { useState } from "react"
import { useMutation } from "@apollo/client"
import { SAVE_ARTICLE_MUTATION } from "../graphql/saveArticle"

export default function UrlSubmitForm() {
  const [url, setUrl] = useState("")
  const [saveArticle, { loading, error, data }] = useMutation(SAVE_ARTICLE_MUTATION)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!url) return
    try {
      await saveArticle({ variables: { url } })
      setUrl("")
    } catch (err) {
      console.error(err)
    }
  }

  return (
    <div className="max-w-xl mx-auto mt-8 p-4 bg-white shadow rounded">
      <form onSubmit={handleSubmit} className="space-y-4">
        <input
          type="url"
          className="w-full border border-gray-300 p-2 rounded"
          placeholder="Enter article URL"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          required
        />
        <button
          type="submit"
          className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
          disabled={loading}
        >
          {loading ? "Submitting..." : "Submit URL"}
        </button>
      </form>

      {error && <p className="text-red-500 mt-2">Error: {error.message}</p>}
      {data && (
        <div className="mt-4">
          <h2 className="font-bold text-gray-700">Summary:</h2>
          <p className="text-gray-800 mt-1">{data.saveArticle.summary}</p>
        </div>
      )}
    </div>
  )
}
