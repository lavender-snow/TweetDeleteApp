document.addEventListener('DOMContentLoaded', function(e) {
	const deleteButton = document.getElementById('delete-button');
	const tweetChecks = document.querySelectorAll('input[name="tweetid"]');
	const tweetCountLabel = document.getElementById('tweet-count-label');

	// 削除ボタン非活性化
	deleteButton.disabled = true;

	// チェックボックスにイベント追加
	[...tweetChecks].forEach(tweetCheck=>{
		tweetCheck.addEventListener('change', function() {

			let count = 0;

			for (tweetCheck of tweetChecks ){
				if (tweetCheck.checked){
					count++;
				}
			}

			tweetCountLabel.innerHTML = count;

			// 1件以上選択されている場合削除ボタンを活性化
			if (count >= 1){
				deleteButton.disabled = false;
			} else {
				deleteButton.disabled = true;
			}

		}, false);
	});

}, false)