init().catch((error) => {
    console.error(error);
});

async function init() {
    await attachTilesInputHandler();
    await attachBoardInputHandler();
}

async function attachTilesInputHandler() {
    for (const input of document.querySelectorAll(".your-tiles__input")) {
        input.addEventListener("keyup", async () => {
            await handleTileChange();
        });
    }
}

async function handleTileChange() {
    const tiles = [];

    for (const input of document.querySelectorAll(".your-tiles__input")) {
        tiles.push({
            id: input.id,
            value: input.value,
        });
    }

    await saveTiles(tiles);
}

async function saveTiles(tiles) {
    const response = await fetch("/tiles", {
        method: "POST",
        body: JSON.stringify({ tiles }),
        headers: {
            "Content-Type": "application/json",
        },
    });

    return await response.json();
}

async function attachBoardInputHandler() {
    for (const input of document.querySelectorAll("input")) {
        input.addEventListener("keyup", async () => {
            await handleBoardChange();
        });
    }
}

async function handleBoardChange() {
    const board = [];

    for (const row of document.querySelectorAll(".current-board__row")) {
        const boardRow = [];
        for (const tile of row.querySelectorAll(".current-board__input")) {
            boardRow.push({
                id: tile.id,
                value: tile.value,
            });
        }
        board.push(boardRow);
    }

    await saveBoard(board);
}

async function saveBoard(board) {
    const response = await fetch("/board", {
        method: "POST",
        body: JSON.stringify({ board }),
        headers: {
            "Content-Type": "application/json",
        },
    });

    return await response.json();
}
