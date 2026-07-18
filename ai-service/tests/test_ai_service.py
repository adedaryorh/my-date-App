"""Fast unit tests for ranking logic; models and PostgreSQL are integration concerns."""
from app.scoring import rank_candidates, score_candidate


def test_shared_interests_boost_similarity():
    assert abs(score_candidate(0.50, shared_interests=2) - 0.60) < 1e-9


def test_positive_and_negative_feedback_adjust_score():
    assert abs(score_candidate(0.5, interactions={"like": 2, "skip": 1}) - 0.55) < 1e-9
    assert abs(score_candidate(0.5, interactions={"report": 1}) - 0.25) < 1e-9


def test_interaction_count_is_capped():
    assert score_candidate(0, interactions={"follow": 20}) == 1.0


def test_rank_candidates_descending_and_limited():
    candidates = [{"id": "low", "score": 0.1}, {"id": "high", "score": 0.9}, {"id": "mid", "score": 0.5}]
    assert [item["id"] for item in rank_candidates(candidates, 2)] == ["high", "mid"]
    assert candidates[0]["id"] == "low"
