yay


				$ajax_nonce = wp_create_nonce( 'nonce' )

				add_action( 'wp_ajax_', '' );
				function (){
					check_ajax_referer( 'nonce', 'nonce' );
					// do some work
					wp_send_json_success();
				}

				

				$ajax_nonce = wp_create_nonce( 'nonce' )

				add_action( 'wp_ajax_', '' );
				function (){
					check_ajax_referer( 'nonce', 'nonce' );
					// do some work
					wp_send_json_success();
				}

				

				var data = {
					'nonce': nonce,
					'action': '%!s(BADINDEX)',
				};

				$.ajax({
					method: 'POST',
					dataType: 'json',
					url: ajax_url,
					data: data,
					success: function (response) {
						// do some work
					},
					error: function () {
						// something went wrong
					}
				}).always(function () {
					// catch all behavior
				});

				

				var data = {
					'nonce': nonce,
					'action': 'actionstuff',
				};

				$.ajax({
					method: 'POST',
					dataType: 'json',
					url: ajax_url,
					data: data,
					success: function (response) {
						// do some work
					},
					error: function () {
						// something went wrong
					}
				}).always(function () {
					// catch all behavior
				});

				