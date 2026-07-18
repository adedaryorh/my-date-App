"""Pure, dependency-free recommendation scoring helpers."""
from typing import Any, Dict, List, Optional

INTERACTION_WEIGHTS = {"follow": 0.2, "like": 0.1, "skip": -0.15, "report": -0.25}


def score_candidate(base_similarity: float, shared_interests: int = 0,
                    interactions: Optional[Dict[str, int]] = None) -> float:
    score = float(base_similarity) + (0.05 * max(0, shared_interests))
    for action, count in (interactions or {}).items():
        score += INTERACTION_WEIGHTS.get(action, 0.0) * min(max(int(count), 0), 5)
    return score


def rank_candidates(candidates: List[Dict[str, Any]], limit: int) -> List[Dict[str, Any]]:
    return sorted(candidates, key=lambda item: item.get("score", 0.0), reverse=True)[:limit]
