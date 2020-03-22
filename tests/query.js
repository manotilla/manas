db.processes.aggregate([
    { $group: {
            _id: { firstField: "$ipc", secondField: "$ipc" },
            uniqueIds: { $addToSet: "$cmd" },
            count: { $sum: 1 }
        }},
    { $match: {
            count: { $gt: 1 }
        }}
])