var interval = null;

function initialize() {
    showToast();
    sync();
    interval = setInterval(sync, 1000);
}

function sync() {
    syncTimer();
    syncEvents();
    syncVotes();
    syncResults();
}

function syncTimer() {

    var start = new Date(room.resetTs);
    var now = new Date();
    var diff = now - start;
    var mins = Math.floor(diff / 60000);
    var secs = Math.floor((diff/1000)%60);
    var timer = padZero(""+mins)+":"+padZero(""+secs);
    var elem = document.getElementById("timer");
    elem.innerHTML = timer;

    function padZero(s) {
        if (s.length < 2) {
            return "0" + s;
        }
        return s;
    }
}

function syncEvents() {
    fetch("/rooms/" + room.name + "/events")
        .then(r => r.json())
        .then(d => {
            if (shouldReveal(d) || shouldReset(d)) {
                reload();
            }
        });

        function shouldReveal(d) {
            return d.revealed && !room.revealed;
        }

        function shouldReset(d) {
            return d.resetTs > room.resetTs
        }
}

function syncVotes() {
    fetch("/rooms/" + room.name + "/userlist")
        .then(r => r.text())
        .then(s => {
            var el = document.getElementById("userlist");
            el.innerHTML = s;
        });
}

function syncResults() {
    if (room.revealed) {
        fetch("/rooms/" + room.name + "/results")
            .then(r => r.text())
            .then(s => {
                var el = document.getElementById("results");
                el.innerHTML = s;
            });
    }
}

function vote(v) {
    fetch("/rooms/" + room.name + "/vote", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: v
    });
}

function showToast() {

    if (room.revealed) {
        makeToast(room.revealedBy, "revealed the votes");
    } else if (room.resetBy) {
        makeToast(room.resetBy, "started a new story")
    }


    function makeToast(name, action) {
        M.toast({
            html: "<span class='amber-text'>"
                + name + "</span>&nbsp;" + action,
            displayLength: 15000,
         });
    }

}

async function reveal() {
    await fetch("/rooms/" + room.name + "/reveal", {
        method: "POST"
    });
    reload();
}

async function reset() {
    await fetch("/rooms/" + room.name + "/reset", {
        method: "POST"
    });
    reload();
}

function reload() {
    clearInterval(interval);
    window.location.reload(true);
}

(() => {
    const radios = new Map();
    document.querySelectorAll('input[name=votes][accesskey]').forEach(e => {
        const accesskey = e.getAttribute('accesskey');
        radios.set(accesskey, e);
    });

    document.addEventListener('keyup', e => radios.get(e.key)?.click());
})();
