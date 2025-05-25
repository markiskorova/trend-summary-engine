import UrlSubmitForm from "./components/UrlSubmitForm"

function App() {
  return (
    <div className="min-h-screen bg-gray-100 p-6">
      <h1 className="text-3xl font-bold text-center text-gray-800">Trend Summary Engine</h1>
      <p className="text-center mt-2 text-gray-600">Submit an article URL to get a summary.</p>
      <UrlSubmitForm />
    </div>
  )
}

export default App
