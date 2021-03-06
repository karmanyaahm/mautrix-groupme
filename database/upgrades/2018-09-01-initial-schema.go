package upgrades

import (
	"gorm.io/gorm"
)

func init() {
	upgrades[0] = upgrade{"Initial schema", func(tx *gorm.DB, ctx context) error {
		// _, err := tx.Exec(`CREATE TABLE IF NOT EXISTS portal (
		// 	jid      VARCHAR(255),
		// 	receiver VARCHAR(255),
		// 	mxid     VARCHAR(255) UNIQUE,

		// 	name   VARCHAR(255) NOT NULL,
		// 	topic  VARCHAR(512) NOT NULL,
		// 	avatar VARCHAR(255) NOT NULL,
		// 	avatar_url VARCHAR(255),
		// 	encrypted BOOLEAN NOT NULL DEFAULT false

		// 	PRIMARY KEY (jid, receiver)
		// )`)
		// if err != nil {
		// 	return err
		// }

		// _, err = tx.Exec(`CREATE TABLE IF NOT EXISTS puppet (
		// 	jid          VARCHAR(255) PRIMARY KEY,
		// 	avatar       VARCHAR(255),
		// 	displayname  VARCHAR(255),
		// 	name_quality SMALLINT,
		// 	custom_mxid VARCHAR(255),
		// 	access_token VARCHAR(1023),
		// 	next_batch VARCHAR(255),
		// 	avatar_url VARCHAR(255),
		// 	enable_presence BOOLEAN NOT NULL DEFAULT true,
		// 	enable_receipts BOOLEAN NOT NULL DEFAULT true
		// )`)
		// if err != nil {
		// 	return err
		// }

		// _, err = tx.Exec(`CREATE TABLE IF NOT EXISTS "user" (
		// 	mxid VARCHAR(255) PRIMARY KEY,
		// 	jid  VARCHAR(255) UNIQUE,

		// 	management_room VARCHAR(255),

		// 	client_id    VARCHAR(255),
		// 	client_token VARCHAR(255),
		// 	server_token VARCHAR(255),
		// 	enc_key      bytea,
		// 	mac_key      bytea,
		// 	last_connection BIGINT NOT NULL DEFAULT 0
		// )`)
		// if err != nil {
		// 	return err
		// }

		// _, err = tx.Exec(`CREATE TABLE IF NOT EXISTS "user_portal" (
		// 	user_jid        VARCHAR(255),
		// 	portal_jid      VARCHAR(255),
		// 	portal_receiver VARCHAR(255),
		// 	in_community BOOLEAN NOT NULL DEFAULT FALSE,

		// 	PRIMARY KEY (user_jid, portal_jid, portal_receiver),

		// 	FOREIGN KEY (user_jid) REFERENCES "user"(jid) ON DELETE CASCADE,
		// 	FOREIGN KEY (portal_jid, portal_receiver) REFERENCES portal(jid, receiver) ON DELETE CASCADE
		// )`)
		// if err != nil {
		// 	return err
		// }

		// _, err = tx.Exec(`CREATE TABLE IF NOT EXISTS message (
		// 	chat_jid      VARCHAR(255),
		// 	chat_receiver VARCHAR(255),
		// 	jid           VARCHAR(255),
		// 	mxid          VARCHAR(255) NOT NULL UNIQUE,
		// 	sender        VARCHAR(255) NOT NULL,
		// 	content       bytea        NOT NULL,
		// 	timestamp 	  BIGINT       NOT NULL DEFAULT 0,

		// 	PRIMARY KEY (chat_jid, chat_receiver, jid),
		// 	FOREIGN KEY (chat_jid, chat_receiver) REFERENCES portal(jid, receiver) ON DELETE CASCADE
		// )`)
		// if err != nil {
		// 	return err
		// }

		// _, err = tx.Exec(`CREATE TABLE IF NOT EXISTS mx_registrations (
		// 	user_id VARCHAR(255) PRIMARY KEY
		// )`)
		// if err != nil {
		// 	return err
		// }

		// _, err = tx.Exec(`CREATE TABLE IF NOT EXISTS mx_room_state (
		// 	room_id      VARCHAR(255) PRIMARY KEY,
		// 	power_levels TEXT
		// )`)
		// if err != nil {
		// 	return err
		// }

		// _, err = tx.Exec(`CREATE TABLE IF NOT EXISTS mx_user_profile (
		// 	room_id     VARCHAR(255),
		// 	user_id     VARCHAR(255),
		// 	membership  VARCHAR(15) NOT NULL,
		// 	PRIMARY KEY (room_id, user_id),
		// 	displayname TEXT,
		// 	avatar_url VARCHAR(255)
		// )`)
		// if err != nil {
		// 	return err
		// }

		// _, err = tx.Exec(`CREATE TABLE IF NOT EXISTS crypto_olm_session (
		// 	session_id   CHAR(43)  PRIMARY KEY,
		// 	sender_key   CHAR(43)  NOT NULL,
		// 	session      bytea     NOT NULL,
		// 	created_at   timestamp NOT NULL,
		// 	last_used    timestamp NOT NULL
		// )`)
		// if err != nil {
		// 	return err
		// }

		// _, err = tx.Exec(`CREATE TABLE crypto_megolm_inbound_session (
		// 	session_id   CHAR(43)     PRIMARY KEY,
		// 	sender_key   CHAR(43)     NOT NULL,
		// 	signing_key  CHAR(43)     NOT NULL,
		// 	room_id      VARCHAR(255) NOT NULL,
		// 	session      bytea        NOT NULL,
		// 	forwarding_chains bytea   NOT NULL
		// )`)
		// if err != nil {
		// 	return err
		// }

		// _, err = tx.Exec(`CREATE TABLE crypto_device (
		// 	user_id      VARCHAR(255),
		// 	device_id    VARCHAR(255),
		// 	identity_key CHAR(43)      NOT NULL,
		// 	signing_key  CHAR(43)      NOT NULL,
		// 	trust        SMALLINT      NOT NULL,
		// 	deleted      BOOLEAN       NOT NULL,
		// 	name         VARCHAR(255)  NOT NULL,

		// 	PRIMARY KEY (user_id, device_id)
		// )`)

		// _, err = tx.Exec(`CREATE TABLE crypto_tracked_user (
		// 	user_id VARCHAR(255) PRIMARY KEY
		// )`)
		// if err != nil {
		// 	return err
		// }

		// _, err = tx.Exec(`CREATE TABLE crypto_message_index (
		// 	sender_key CHAR(43),
		// 	session_id CHAR(43),
		// 	"index"    INTEGER,
		// 	event_id   VARCHAR(255) NOT NULL,
		// 	timestamp  BIGINT       NOT NULL,

		// 	PRIMARY KEY (sender_key, session_id, "index")
		// )`)
		// if err != nil {
		// 	return err
		// }

		// _, err = tx.Exec(`CREATE TABLE crypto_account (
		// 	device_id  VARCHAR(255) PRIMARY KEY,
		// 	shared     BOOLEAN      NOT NULL,
		// 	sync_token TEXT         NOT NULL,
		// 	account    bytea        NOT NULL
		// )`)
		// if err != nil {
		// 	return err
		// }

		// _, err = tx.Exec(`CREATE TABLE crypto_megolm_outbound_session (
		// 	room_id       VARCHAR(255) PRIMARY KEY,
		// 	session_id    CHAR(43)     NOT NULL UNIQUE,
		// 	session       bytea        NOT NULL,
		// 	shared        BOOLEAN      NOT NULL,
		// 	max_messages  INTEGER      NOT NULL,
		// 	message_count INTEGER      NOT NULL,
		// 	max_age       BIGINT       NOT NULL,
		// 	created_at    timestamp    NOT NULL,
		// 	last_used     timestamp    NOT NULL
		// )`)
		// if err != nil {
		// 	return err
		// }

		return nil
	}}
}
