import React, { useEffect, useState } from "react";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { motion, AnimatePresence } from "framer-motion";

/**
 * FunFactsApp – minimalistic responsive frontend for the History Fun‑Facts API
 *
 * Expected backend:
 *   GET  /   → { status: "200 OK", info: "some fact" }
 *   POST /   → { status: "200 OK", id: 10 }  OR error format

export default function FunFactsApp() {
    const [fact, setFact] = useState(null);
    const [input, setInput] = useState("");
    const [message, setMessage] = useState(null);
    const [loading, setLoading] = useState(false);

    // Fetch a random fact from the API
    const fetchRandomFact = async (link) => {
        try {
            const res = await fetch(link);
            const data = await res.json();
            if (data.info) setFact(data.info);
        } catch (err) {
            console.error(err);
        }
    };

    useEffect(() => {
        fetchRandomFact("https://localhost:8080/fuckingnigger1");
    }, []);

    // Submit a new fact to the API
    const submitFact = async (e) => {
        e.preventDefault();
        if (!input.trim()) return;
        setLoading(true);
        try {
            const res = await fetch("https://localhost:4000/fuckingnigger2", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ fact: input.trim() }),
            });
            const data = await res.json();
            if (res.ok) {
                setMessage({ type: "success", text: `Thanks! Saved with id #${data.id}` });
                setInput("");
            } else {
                setMessage({ type: "error", text: data.error || "Something went wrong" });
            }
        } catch (err) {
            setMessage({ type: "error", text: "Network error" });
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-100 to-slate-300 p-4">
            <div className="w-full max-w-md space-y-6">
                {/* Display random fact }
                <Card className="shadow-xl">
                    <CardContent className="p-6">
                        <h1 className="text-2xl font-bold mb-4 text-center">Random History Fact</h1>
                        <AnimatePresence mode="wait">
                            {fact && (
                                <motion.p
                                    key={fact}
                                    initial={{ opacity: 0, y: 10 }}
                                    animate={{ opacity: 1, y: 0 }}
                                    exit={{ opacity: 0, y: -10 }}
                                    transition={{ duration: 0.3 }}
                                    className="text-center text-lg"
                                >
                                    {fact}
                                </motion.p>
                            )}
                        </AnimatePresence>
                        <div className="flex justify-center mt-4">
                            <Button onClick={fetchRandomFact} variant="outline">
                                New Fact
                            </Button>
                        </div>
                    </CardContent>
                </Card>

                {/* Submit new fact }
                <Card className="shadow-xl">
                    <CardContent className="p-6">
                        <h2 className="text-xl font-semibold mb-4 text-center">Submit Your Own Fact</h2>
                        <form onSubmit={submitFact} className="space-y-4">
              <textarea
                  className="w-full p-3 rounded-xl border focus:outline-none focus:ring focus:ring-sky-400"
                  rows={4}
                  placeholder="Did you know...?"
                  value={input}
                  onChange={(e) => setInput(e.target.value)}
              />
                            <motion.div whileHover={{ scale: 1.05 }} whileTap={{ scale: 0.95 }}>
                                <Button className="w-full" disabled={loading} type="submit">
                                    {loading ? "Submitting..." : "Submit Fact"}
                                </Button>
                            </motion.div>
                        </form>
                        {message && (
                            <motion.p
                                initial={{ opacity: 0 }}
                                animate={{ opacity: 1 }}
                                className={`mt-4 text-center ${message.type === "error" ? "text-red-600" : "text-green-600"}`}
                            >
                                {message.text}
                            </motion.p>
                        )}
                    </CardContent>
                </Card>
            </div>
        </div>
    );
}

*/