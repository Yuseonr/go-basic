package main

// 21 March 2026
// CH - 15 : Enums

// ========================================== L1 : Lack of Enums ==========================================
import ("fmt")

func (a *analytics) handleEmailBounce(em email) error {
	err := em.recipient.updateStatus(em.status)
	if err != nil {return fmt.Errorf("error updating user status: %w", err)}

	err = a.track(em.status)
	if err != nil {return fmt.Errorf("error tracking user bounce: %w", err)}

	return nil
}

// ===================

type email struct {
	status    string
	recipient *user
}

type user struct {
	email  string
	status string
}

type analytics struct {
	totalBounces int
}

func (u *user) updateStatus(status string) error {
	if status != "email_bounced" && status != "email_failed" {
		return fmt.Errorf("invalid status: %s", status)
	}
	u.status = status
	return nil
}

func (a *analytics) track(event string) error {
	if event != "email_bounced" {
		return fmt.Errorf("invalid event: %s", event)
	}
	a.totalBounces++
	return nil
}

// ========================================== L2 : Type Definitions ==========================================

// checkPermission(User) salah -> checkPermission(Admin)

// ========================================== L3 : Type Definitions ==========================================

// No

// ========================================== L4 : Iota ==========================================

type emailStatus int

const (
	EmailBounced emailStatus = iota
	EmailInvalid
	EmailDelivered
	EmailOpened
)